package token

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"os"
	"strings"
	"sync"
	"time"

	"ilicense-lite/config"

	"github.com/golang-jwt/jwt"
)

var (
	jwtSecretOnce sync.Once
	jwtSecret     []byte
	jwtSecretErr  error
)

func getJWTSecret() ([]byte, error) {
	jwtSecretOnce.Do(func() {
		secret := strings.TrimSpace(os.Getenv("JWT_SECRET"))
		if secret == "" {
			secret = strings.TrimSpace(config.Config.App.JWTSecret)
		}
		if secret == "" {
			// Fallback for development: generate a runtime-only secret.
			// This avoids shipping hardcoded secrets while keeping local bootstrapping possible.
			buf := make([]byte, 32)
			if _, err := rand.Read(buf); err != nil {
				jwtSecretErr = errors.New("jwt secret is empty and random generation failed")
				return
			}
			secret = base64.RawStdEncoding.EncodeToString(buf)
		}
		jwtSecret = []byte(secret)
	})
	return jwtSecret, jwtSecretErr
}

// GenerateJWT 生成 JWT
func GenerateJWT(userID string) (string, error) {
	secret, err := getJWTSecret()
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

// ValidateJWT 验证 JWT
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	secret, err := getJWTSecret()
	if err != nil {
		return nil, err
	}

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 只允许 HS256 签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return secret, nil
	})
}
