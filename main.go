package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"./handlers"
	"./models"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	models.InitAwsServices()

	router := httprouter.New()
	router.POST("/reservations/:reservationId/email-confirmation", handlers.CreateTicketRelease)

	fmt.Println("Server started on port " + os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}
