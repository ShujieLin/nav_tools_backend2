package config

import "github.com/jinzhu/gorm"

var Cfg *Config

type Config struct {
	DbPath     string
	ConfigPath string // 新增配置路径
}

func LoadConfig() {
	Cfg = &Config{
		DbPath:     "./navtools.db",
		ConfigPath: "./configs", // 默认配置目录
	}
}

func AutoMigrate(db *gorm.DB) {
	// 自动迁移模型将在此处添加
}
