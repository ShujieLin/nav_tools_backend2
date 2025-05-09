package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"nav_tools_backend2/models"
	"os"

	"github.com/jinzhu/gorm"
)

type LinkData struct {
	Category  string `json:"category"`
	UrlTitle  string `json:"url_title"`
	UrlDes    string `json:"url_des"`
	DataUrl   string `json:"data_url"`
	InnerLink string `json:"inner_link"`
}

func ImportLinksFromJSON(db *gorm.DB, filePath string) error {
	// 读取JSON文件
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var links []LinkData
	if err := json.Unmarshal(byteValue, &links); err != nil {
		return err
	}

	// 批量导入数据
	for _, link := range links {
		newLink := models.Link{
			Category:  link.Category,
			UrlTitle:  link.UrlTitle,
			UrlDes:    link.UrlDes,
			DataUrl:   link.DataUrl,
			InnerLink: link.InnerLink,
		}
		if err := db.Create(&newLink).Error; err != nil {
			log.Printf("导入失败: %v", err)
			continue
		}
	}

	return nil
}
