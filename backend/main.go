package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/Aeswolf/equipment-database-management/api"
	"github.com/Aeswolf/equipment-database-management/database"
	"github.com/Aeswolf/equipment-database-management/server"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		logError(err)
	}

	dsn := os.Getenv("DATA_SOURCE_NAME")

	store, err := database.New(dsn)

	if err != nil {
		logError(err)
	}

	apiServer := api.NewServer("4000", store)

	if err := server.Run(apiServer); err != nil {
		logError(err)
	}
}

func logError(err error) {
	log.Fatalf("%+v\n", err)
}
