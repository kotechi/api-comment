package model

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Text   string 
	UserID uint
	PostID uint
}

func main() {
	dsn := "host=localhost user=postgres password=@pasific12op dbname=kubikitdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(&Comment{}); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration successful!")
}
