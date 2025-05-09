package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"uniqueIndex;not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	FirstName string
	LastName  string
}

// 其他模型结构体可在此继续添加
