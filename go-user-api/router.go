package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func newRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	handler := newUserHandler(db)

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
