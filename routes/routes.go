package routes

import (
	"nav_tools_backend2/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RegisterRoutes(r *gin.Engine) {
	// 健康检查接口
	r.GET("/health", healthCheck)

	// API v1路由组
	v1 := r.Group("/api/v1")
	{
		v1.POST("/users", createUser)
		v1.GET("/users", listUsers)
		v1.GET("/users/:id", getUser)
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}

func createUser(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Create(&user)
	c.JSON(201, user)
}

func listUsers(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)
	var users []models.User
	db.Find(&users)
	c.JSON(200, users)
}

func getUser(c *gin.Context) {
	db := c.MustGet("DB").(*gorm.DB)
	id := c.Param("id")
	var user models.User
	if db.First(&user, id).RecordNotFound() {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}
	c.JSON(200, user)
}
