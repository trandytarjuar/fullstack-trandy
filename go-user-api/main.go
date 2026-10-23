package main

import (
	"log"
	"os"
)

const (
	defaultAddress = ":8081"
	defaultDBPath  = "users.db"
)

// @title User API
// @version 1.0
// @description REST API sederhana untuk mengelola data user.
// @host localhost:8081
// @BasePath /api
func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = defaultDBPath
	}

	db, err := openDatabase(dbPath)
	if err != nil {
		log.Fatalf("initialize database: %v", err)
	}

	router := newRouter(db)
	if err := router.Run(defaultAddress); err != nil {
		log.Fatalf("run server: %v", err)
	}
}
