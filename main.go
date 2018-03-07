package main

import (
	"fmt"
	"log"
	"os"

	"./routes"
	"./db"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()

	fmt.Println("Server started on port " + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), routes.BuildRouter())
}
