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

	err = db.AutoMigrate(&Ansprechpartner{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Lieferant{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Kanban{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Post{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Mitarbeiter{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	return db
}
