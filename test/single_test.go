package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Singleton 结构体代表单例
type Singleton struct {
}

// instance 保存Singleton的唯一实例
var instance *Singleton

// once 确保singletonInstance函数只执行一次
var once sync.Once

// GetInstance 是获取Singleton实例的方法
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func (s *Singleton) SayHello() {
	fmt.Println("Hello from the Singleton!")
}

func TestSingle(t *testing.T) {
	s1 := GetInstance()
	s2 := GetInstance()

	// 检查两个实例是否相同，以验证单例模式是否正确实现
	if s1 == s2 {
		fmt.Println("s1 and s2 are the same instance.")
	} else {
		fmt.Println("s1 and s2 are different instances.")
	}

	s1.SayHello()
}

var (
	a  = 1
	wg = sync.WaitGroup{}
)

func TestFlightGroup_Do(t *testing.T) {
	single := Single
	for i := 0; i < 3; i++ {
		go func() {
			do, err := single.Do("key", func() (interface{}, error) {
				return addOne()
			})
			fmt.Println(do, err)
		}()
	}
	time.Sleep(time.Second * 10)
}

func addOne() (int, error) {
	a += 1
	time.Sleep(1 * time.Second)
	return a, nil
}

// panic 不能为负数
func TestDoneOver(t *testing.T) {
	wg.Add(1)
	wg.Done()
	wg.Done()
	wg.Wait()
}
