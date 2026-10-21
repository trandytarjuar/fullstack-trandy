package main

import (
	"log"
	"os"
)

const (
	defaultAddress = ":8081"
	defaultDBPath  = "users.db"
)

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
