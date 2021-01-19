package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


func main() {
	a := App{}
	
	err := godotenv.Load(os.ExpandEnv(".env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	a.Initialize(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
        os.Getenv("DB_DRIVER"))

	a.initializeRoutes()
	a.Run(":8080")
}