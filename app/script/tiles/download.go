// Package main
// @Description: 从openstreetmap下载地图瓦片数据
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	tileURL = "https://tile.tracestrack.com/topo__/%d/%d/%d.png?key=%s"
	// 瓦片存储目录
	outputDir  = "/Users/ibingbo/go/src/gitlab.com/gin-app/data/tiles"
	minZoom    = 18
	MaxZoom    = 19
	maxRetries = 5 // 设置最大重试次数
)

// API Key 列表
var (
	keys = []string{
		// "130cab5ea1682bb3161fa926c80542eb",
		// "3adf32fe2c68504e35f4c5ab0cefa38d",
		// "c330b644695e73b0b4139aa6cef250d9",
		// "659330e0a7eda5a909d9bb81c72032a0",
		// "def91a6870dc5dd918fac0971890c0cc",
		// "a6a3b785991464fafeccd358d68e597e",
		// "7dde9ea64e26d2fd2293f5165df91f65",
		// "98148343716b8fedcdb92309eb0ea4bf",
		// "7aeb516a226072a500cd957a715a61ad",
		// ibingbo@163.com
		"2696a4acc97d3a4e836e2a918ef7ea79",
		"1a825d984b777a38331e48bb7f8176f0",
		// zhangbingbingx@163.com
		"7cddb1b36ee2621de60be6e5e3992653",
		"b7bce4c6d0029be8ecfa948f75b65775",
		// bingbo.zh@gmail.com
		"17bec8a281324d82b8c3419ffdacee74",
		"4aa096c09410dc0348bf86bae727ad65",
		// ibingbo.zh@outlook.com
		"e2b9fb73aa4981f41a55022cee0ff015",
		"aeffe924d88b663e7bcbb699d3bd4495",
		// mingming qq
		"7d55c7a5dfebd77fd8c22ff4143d0a11",
		"fc9d00f2c966b86c2aaa09cdd2aeb776",
	}
	// 失败任务队列
	failedChan = make(chan Tile, 1000)
	done       bool
)

type (
	Tile struct {
		z, x, y, retry int
	}
)

func main() {
	rand.Seed(time.Now().UnixNano()) // 设置随机种子
	// 丽水市的经纬度范围
	minLon, maxLat := 118.66333, 28.97511 // 左上角（西北）
	maxLon, minLat := 120.46509, 27.34667 // 右下角（东南）

	// 可以根据这个范围批量下载瓦片
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 5) // 限制并发数为 10

	// 启动重试
	wg.Add(1)
	go func() {
		defer wg.Done()
		lastedTime := time.Now().Unix()
		for {
			select {
			case tile, ok := <-failedChan:
				if !ok {
					log.Println("失败队列已关闭")
					return
				}
				log.Printf("收到失败任务: %v\n", tile)
				if !downloadTile(tile.z, tile.x, tile.y) {
					failedChan <- tile
					if tile.retry < maxRetries {
						failedChan <- Tile{z: tile.z, x: tile.x, y: tile.y, retry: tile.retry + 1}
					} else {
						log.Printf("瓦片下载失败次数过多，放弃: %v\n", tile)
					}
				}
				lastedTime = time.Now().Unix()
			default:
				if done && len(failedChan) == 0 && (time.Now().Unix()-lastedTime > 300) {
					log.Println("超过5分钟没有新任务，认为所有任务已完成")
					close(failedChan)
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// 开始下载
	for zoom := minZoom; zoom <= MaxZoom; zoom++ {
		// 计算瓦片坐标范围
		minX, minY := lonLatToTile(minLon, maxLat, zoom) // 左上角
		maxX, maxY := lonLatToTile(maxLon, minLat, zoom) // 右下角
		log.Printf("zoom %d - x 范围: %d - %d, y 范围: %d - %d\n", zoom, minX, maxX, minY, maxY)

		// 遍历范围内的所有瓦片
		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				zCopy, xCopy, yCopy := zoom, x, y // 创建新的局部变量，避免变量污染
				wg.Add(1)
				semaphore <- struct{}{} // 控制并发
				go func(z, x, y int) {
					defer wg.Done()
					defer func() { <-semaphore }()
					if !downloadTile(z, x, y) {
						failedChan <- Tile{z, x, y, 0}
					}
				}(zCopy, xCopy, yCopy)
				time.Sleep(50 * time.Millisecond) // 控制请求频率，避免封禁
			}
		}
	}
	done = true
	wg.Wait()
	log.Println("所有瓦片下载完成！")
}

// 随机获取一个 API Key
func getRandomKey() string {
	return keys[rand.Intn(len(keys))]
}

// 转换经纬度到像素坐标
func lonLatToTile(lon, lat float64, zoom int) (x, y int) {
	n := math.Pow(2, float64(zoom))
	x = int((lon + 180.0) / 360.0 * n)
	latRad := lat * math.Pi / 180.0
	y = int((1 - math.Log(math.Tan(latRad)+(1/math.Cos(latRad)))/math.Pi) / 2 * n)
	return
}

// 瓦片下载器
func downloadTile(z, x, y int) bool {

	key := getRandomKey()
	// 生成瓦片 URL
	url := fmt.Sprintf(tileURL, z, x, y, key)

	// 构造本地存储路径
	tilePath := fmt.Sprintf("%s/%d/%d", outputDir, z, x)
	filePath := fmt.Sprintf("%s/%d.png", tilePath, y)
	tmpPath := filePath + ".tmp" // 临时文件，防止下载失败

	// 检查文件是否已存在，避免重复下载
	if _, err := os.Stat(filePath); err == nil {
		log.Println("已存在，跳过：", filePath)
		return true
	}

	// 发送请求下载瓦片
	resp, err := http.Get(url)
	if err != nil {
		log.Println("下载失败:", err)
		return false
	}
	defer resp.Body.Close()

	// 确保响应是 200 OK
	if resp.StatusCode != http.StatusOK {
		log.Println("无效响应:", resp.StatusCode, url)
		return false
	}

	// 确保目录存在
	if err := os.MkdirAll(tilePath, os.ModePerm); err != nil {
		log.Println("创建目录失败:", err)
		return false
	}

	// 保存瓦片到本地
	outTmpFile, err := os.Create(tmpPath)
	if err != nil {
		log.Println("文件创建失败:", err)
		return false
	}
	defer outTmpFile.Close()

	_, err = io.Copy(outTmpFile, resp.Body)
	if err != nil {
		log.Println("写入文件失败:", err)
		return false
	}

	// 重命名为正式文件
	err = os.Rename(tmpPath, filePath)
	if err != nil {
		log.Printf("重命名失败: %s\n", err)
		return false
	}
	log.Printf("下载完成: %s\n", filePath)
	return true
}
