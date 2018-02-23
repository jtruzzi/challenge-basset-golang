package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"./routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := routes.NewRouter()
	fmt.Println("Server started on http://localhost:5000")
	server := http.ListenAndServe(":" + os.Getenv("PORT"), router)

	log.Fatal(server)
}

