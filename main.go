package main

import (
	"log"
	"nav_tools_backend2/config"
	"nav_tools_backend2/database"
	"nav_tools_backend2/models"
	"nav_tools_backend2/routes"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	log.Println("正在加载配置...")
	config.LoadConfig()
	log.Println("配置加载完成")

	// 连接数据库
	log.Println("正在连接数据库...")
	db, err := database.ConnectSQLite()
	if err != nil {
		log.Fatal("数据库连接失败: ", err)
	}
	defer db.Close()
	log.Println("数据库连接成功")

	// 自动迁移表结构
	db.AutoMigrate(&models.Link{})

	// 检查初始化标记
	type InitFlag struct {
		ID     uint `gorm:"primary_key"`
		Inited bool
	}

	// 自动迁移标记表
	db.AutoMigrate(&InitFlag{}, &models.Link{})

	var flag InitFlag
	db.FirstOrCreate(&flag, InitFlag{ID: 1})

	if !flag.Inited {
		jsonPath := filepath.Join(".", "links2.json")
		log.Println("首次启动，正在导入初始数据...")

		// 先清空现有数据（避免重复）
		if err := db.Exec("DELETE FROM links").Error; err != nil {
			log.Printf("清空links表失败: %v", err)
		}

		// 导入数据
		if err := database.ImportLinksFromJSON(db, jsonPath); err != nil {
			log.Printf("数据导入失败: %v", err)
			return // 失败时直接退出，不设置标记
		}

		// 只有全部导入成功才设置标记
		flag.Inited = true
		if err := db.Save(&flag).Error; err != nil {
			log.Printf("保存初始化标记失败: %v", err)
		}
		log.Println("初始数据导入完成")
	}

	// 初始化Gin引擎
	log.Println("初始化Gin引擎...")
	r := gin.Default()
	log.Println("Gin引擎初始化完成")

	// 注册中间件
	log.Println("注册数据库中间件...")
	r.Use(database.InjectDB(db))

	// 注册路由
	log.Println("注册路由...")
	routes.RegisterRoutes(r)

	// 启动服务
	log.Println("服务启动，监听端口:8080")
	r.Run(":8080")
}
