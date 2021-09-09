package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// 可配置项结构体
type Configuration struct {
	StaticPath string
	Port       string
}

// 配置数据
var Config Configuration

// 加载json文件配置
func LoadConfig(path string) {
	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
	}

	// 创建json解析器
	decoder := json.NewDecoder(file)

	// json解析并将数据填入Configuration中
	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Println(err.Error())
	}
}
