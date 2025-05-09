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

	// 导入JSON数据
	jsonPath := filepath.Join(".", "links2.json") // 使用filepath构造路径
	log.Println("正在导入links2.json数据...path:", jsonPath)
	if err := database.ImportLinksFromJSON(db, jsonPath); err != nil {
		log.Printf("数据导入失败: %v", err)
	} else {
		log.Println("数据导入成功")
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
