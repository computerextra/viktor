package db

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase(connString string) *Database {
	db := initDatabase(connString)
	return &Database{
		db: db,
	}
}

func initDatabase(connString string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(connString), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %v", err))
	}

	db.AutoMigrate(&Ansprechpartner{})
	db.AutoMigrate(&Lieferant{})
	db.AutoMigrate(&Kanban{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Status{})
	db.AutoMigrate(&Mitarbeiter{})
	db.AutoMigrate(&User{})

	return db
}
