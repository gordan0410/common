package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

/*
config 綁定系統環境變數到結構體
*/

type BaseConfig[T any] struct {
	Config *T
}

func NewConfig[T any]() *BaseConfig[T] {
	var baseConfig BaseConfig[T]
	var config T
	baseConfig.Config = &config
	baseConfig.init()
	return &baseConfig
}

// init 初始化配置
func (t *BaseConfig[T]) init() {
	err := env.Parse(t.Config)
	if err != nil {
		log.Panic("Error parsing environment variables: ", err)
	}
}

// InitDotEnv local開發 讀取 .env 檔案
func InitDotEnv(runFlag string, path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Println("Error loading .env file", err)
	}
}
