package models

import "github.com/jinzhu/gorm"

type Link struct {
	gorm.Model
	Category  string `gorm:"type:varchar(255)" json:"category"`
	UrlTitle  string `gorm:"type:varchar(255)" json:"url_title"`
	UrlDes    string `gorm:"type:text" json:"url_des"`
	DataUrl   string `gorm:"type:varchar(512)" json:"data_url"`
	InnerLink string `gorm:"type:varchar(512)" json:"inner_link"`
}
