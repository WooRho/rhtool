package rhtool_common

import (
	"encoding/gob"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

// 假设我们的缓存实体
type CacheItem struct {
	Key   string
	Value interface{}
}

// 缓存结构体，包含一个sync.Map和文件持久化的相关信息
type PersistentCache struct {
	cache     sync.Map
	cacheFile string
}

// 初始化PersistentCache实例
func newPersistentCache(filePath string) (*PersistentCache, error) {
	pc := &PersistentCache{
		cacheFile: filePath,
	}

	// 加载现有的缓存数据
	if err := pc.load(); err != nil {
		// 如果加载失败，不影响后续使用，但应该记录错误信息
		log.Printf("Failed to load cache from file: %v", err)
	}

	return pc, nil
}

// 保存缓存到文件
func (pc *PersistentCache) save() error {
	// 先尝试删除目标文件（如果有）
	if err := os.Remove(pc.cacheFile); err != nil && !os.IsNotExist(err) {
		return err
	}

	// 直接创建目标文件
	file, err := os.Create(pc.cacheFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	cacheItems := make(map[string]CacheItem)

	pc.cache.Range(func(key, value interface{}) bool {
		cacheItems[key.(string)] = CacheItem{Key: key.(string), Value: value}
		return true
	})

	if err := encoder.Encode(cacheItems); err != nil {
		return err
	}

	return nil
}

// 加载文件中的缓存数据
func (pc *PersistentCache) load() error {
	file, err := os.Open(pc.cacheFile)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if os.IsNotExist(err) {
		return nil // 如果文件不存在，忽略错误，视为缓存为空
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	var cacheItems map[string]CacheItem
	if err := decoder.Decode(&cacheItems); err != nil {
		return err
	}

	for key, item := range cacheItems {
		pc.cache.Store(key, item.Value)
	}

	return nil
}

// 存储数据到缓存
func (pc *PersistentCache) Set(key string, value interface{}) {
	pc.cache.Store(key, value)

	// 定期或每次更新缓存时将其写入文件实现持久化
	// 这里仅作示例，实际场景下可能需要异步处理或加锁控制并发
	pc.save()
}

// 从缓存获取数据
func (pc *PersistentCache) Get(key string) (interface{}, bool) {
	value, ok := pc.cache.Load(key)
	return value, ok
}

func Persistence(key string, id snowflake.ID) (str string) {
	var (
		uniqName = ""
	)
	dir, err := os.Getwd()
	if err != nil {
		return
	}

	cacheFilePath := filepath.Join(dir, "rhtool_common/cache_file.data")

	cache, err := newPersistentCache(cacheFilePath)
	if err != nil {
		return
	}

	switch runtime.GOOS {
	case "windows":
		uniqName, err = os.Hostname()
		if err != nil {
			fmt.Println("Failed to get hostname:", err)
		} else {
			fmt.Printf("Hostname: %s\n", uniqName)
		}
	case "linux":
		uniqName, err = readProductUUID()
		if err != nil {
			fmt.Println("Failed to get product UUID:", err)
		} else {
			fmt.Printf("Product UUID: %s\n", uniqName)
		}
	case "darwin":
		fmt.Println("This is macOS.")
	default:
		fmt.Printf("This is an unrecognized operating system: %s\n", runtime.GOOS)
	}

	key += uniqName

	// 设置缓存值
	cache.Set(key, id)

	// 获取缓存值
	value, ok := cache.Get(key)
	if ok {
		fmt.Printf("Value for key 'key1': %v\n", value)
	} else {
		fmt.Println("Key 'key1' not found in cache.")
	}

	// 更新缓存值
	cache.Set(key, id)

	// 再次获取更新后的缓存值
	newValue, ok := cache.Get(key)
	if ok {
		fmt.Printf("Updated value for key 'key1': %v\n", newValue)
	}

	// 清理缓存（可选）
	// cache.Clear() // 如果你的持久化缓存实现了Clear方法的话

	// 程序退出前，确保缓存已持久化到文件
	if err = cache.save(); err != nil {
		log.Printf("Failed to save cache to file before exit: %v", err)
	}
	return
}

// 读取Linux系统中的主板UUID
func readProductUUID() (string, error) {
	uuidFile := "/sys/class/dmi/id/product_uuid"
	data, err := os.ReadFile(uuidFile)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}
