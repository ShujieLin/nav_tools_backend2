package main

import (
	"nav_tools_backend2/config"
	"nav_tools_backend2/database"
	"nav_tools_backend2/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化配置
	config.LoadConfig()

	// 连接数据库
	db, err := database.ConnectSQLite()
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}
	defer db.Close()

	// 初始化Gin引擎
	r := gin.Default()

	// 注册中间件
	r.Use(database.InjectDB(db))

	// 注册路由
	routes.RegisterRoutes(r)

	// 启动服务
	r.Run(":8080")
}
