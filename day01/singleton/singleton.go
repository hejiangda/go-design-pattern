package main

import "fmt"

// Singleton 饿汉式单例
type Singleton struct {
	a string // 测试字段
}

func (s *Singleton) GetA() string {
	return s.a
}

var singleton *Singleton

func init() {
	singleton = &Singleton{a: "test"}
}

// GetInstance 获取实例
func GetInstance() *Singleton {
	return singleton
}

func main() {

	s := GetInstance()
	fmt.Println(s.GetA())
}
