package libs

import (
	"fmt"
	"log"

	"github.com/laluardian/pk-ata-go-final-project/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database")
	}
	db.AutoMigrate(&models.Book{})
	fmt.Println("connected to database")
	return db
}
