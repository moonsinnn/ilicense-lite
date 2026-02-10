package util

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"log"
	"strings"
)

// ⚠️ 主密钥（建议来自 env）
// masterKey := os.Getenv("LICENSE_MASTER_KEY")
const masterKey = "master-key"

func generateKey(masterKey string) ([]byte, error) {
	hash := sha256.Sum256([]byte(masterKey))
	return hash[:16], nil // 取前 16 字节 => AES-128
}
func pkcs5Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func pkcs5UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("invalid padding size")
	}
	unpadding := int(data[length-1])
	return data[:(length - unpadding)], nil
}
func encryptAES_ECB(plain, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plain = pkcs5Padding(plain, block.BlockSize())
	encrypted := make([]byte, len(plain))

	for bs, be := 0, block.BlockSize(); bs < len(plain); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted, nil
}

func decryptAES_ECB(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decrypted := make([]byte, len(ciphertext))
	for bs, be := 0, block.BlockSize(); bs < len(ciphertext); bs, be = bs+block.BlockSize(), be+block.BlockSize() {
		block.Decrypt(decrypted[bs:be], ciphertext[bs:be])
	}

	return pkcs5UnPadding(decrypted)
}

// Encrypt 私钥
func Encrypt(privateKey string) (string, error) {
	key, err := generateKey(masterKey)
	if err != nil {
		return "", err
	}

	encrypted, err := encryptAES_ECB([]byte(privateKey), key)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// Decrypt 私钥
func Decrypt(encryptedPrivateKey string) (string, error) {
	key, err := generateKey(masterKey)
	if err != nil {
		return "", err
	}

	cipherBytes, err := base64.StdEncoding.DecodeString(encryptedPrivateKey)
	if err != nil {
		return "", err
	}

	decrypted, err := decryptAES_ECB(cipherBytes, key)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

func GetPrivateKey(encryptedPrivateKey string) (*rsa.PrivateKey, error) {
	// 1. 解密
	privateKeyBase64, err := Decrypt(encryptedPrivateKey)
	if err != nil {
		return nil, err
	}

	// 2. Base64 解码
	keyBytes, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		return nil, err
	}

	// 3. 解析 PKCS#8
	key, err := x509.ParsePKCS8PrivateKey(keyBytes)
	if err != nil {
		return nil, err
	}

	// 4. 类型断言（关键）
	rsaKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("private key is not RSA")
	}

	return rsaKey, nil
}
func GenerateActivationCode(jsonData string, privateKey *rsa.PrivateKey) (string, error) {

	dataBytes := []byte(jsonData)
	log.Printf("License数据长度: %d bytes", len(dataBytes))

	// 1. RSA-SHA256 签名
	hashed := sha256.Sum256(dataBytes)

	signBytes, err := rsa.SignPKCS1v15(
		rand.Reader,
		privateKey,
		crypto.SHA256,
		hashed[:],
	)
	if err != nil {
		return "", err
	}
	log.Printf("签名长度: %d bytes (应该是256)", len(signBytes))

	// 2. 组合数据
	buf := new(bytes.Buffer)

	// 数据长度（4字节，大端）
	if err := binary.Write(buf, binary.BigEndian, int32(len(dataBytes))); err != nil {
		return "", err
	}

	// 数据
	if _, err := buf.Write(dataBytes); err != nil {
		return "", err
	}

	// 签名长度（4字节）
	if err := binary.Write(buf, binary.BigEndian, int32(len(signBytes))); err != nil {
		return "", err
	}

	// 签名
	if _, err := buf.Write(signBytes); err != nil {
		return "", err
	}

	// 3. Base64 URL-safe（无 padding）
	base64Str := base64.RawURLEncoding.EncodeToString(buf.Bytes())
	log.Printf("Base64编码后长度: %d", len(base64Str))

	// 4. 格式化
	return formatActivationCode(base64Str), nil
}
func formatActivationCode(base64Str string) string {
	var sb strings.Builder

	for i, ch := range base64Str {
		if i > 0 && i%5 == 0 {
			sb.WriteByte(' ')
		}
		sb.WriteRune(ch)
	}
	return sb.String()
}

func LoadPublicKey(publicKeyBase64 string) (*rsa.PublicKey, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		return nil, err
	}

	pub, err := x509.ParsePKIXPublicKey(keyBytes)
	if err != nil {
		return nil, err
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("public key is not RSA")
	}

	return rsaPub, nil
}
func VerifySignature(
	data []byte,
	signature []byte,
	publicKey *rsa.PublicKey,
) error {

	log.Printf("开始验证签名, 数据长度: %d, 签名长度: %d",
		len(data), len(signature))

	hashed := sha256.Sum256(data)

	err := rsa.VerifyPKCS1v15(
		publicKey,
		crypto.SHA256,
		hashed[:],
		signature,
	)

	log.Printf("签名验证结果: %v", err == nil)
	return err
}
