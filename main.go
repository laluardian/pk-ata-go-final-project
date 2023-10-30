package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/laluardian/pk-ata-go-final-project/libs"
	"github.com/laluardian/pk-ata-go-final-project/routes"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading environment variables")
	}

	dsn := os.Getenv("DATA_SOURCE_NAME")
	db := libs.InitDB(dsn)
	e := routes.InitApi(db)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
