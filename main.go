package main

import (
	"log"
	"os"

	"./routes"
	"./db"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	log.SetOutput(os.Stdout)

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()

	log.Println("Server started on port " + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), routes.BuildRouter())
}
