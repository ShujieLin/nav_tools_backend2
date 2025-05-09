package config

import "github.com/jinzhu/gorm"

var Cfg *Config

type Config struct {
	DbPath string
}

func LoadConfig() {
	Cfg = &Config{
		DbPath: "./navtools.db",
	}
}

func AutoMigrate(db *gorm.DB) {
	// 自动迁移模型将在此处添加
}
