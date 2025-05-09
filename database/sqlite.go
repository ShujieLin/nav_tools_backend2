package database

import (
	"nav_tools_backend2/config"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectSQLite() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", config.Cfg.DbPath)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return db, nil
}

func InjectDB(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}
