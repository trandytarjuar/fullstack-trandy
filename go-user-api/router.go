package main

import (
	_ "example.com/meditap/go-user-api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func newRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	handler := newUserHandler(db)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users := router.Group("/api/users")
	{
		users.GET("", handler.list)
		users.GET("/:userId", handler.get)
		users.POST("", handler.create)
		users.PUT("/:userId", handler.update)
		users.DELETE("/:userId", handler.delete)
	}

	return router
}
