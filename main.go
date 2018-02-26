package main

import (
	"fmt"
	"log"
	"github.com/julienschmidt/httprouter"
	"./handlers"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := httprouter.New()

	router.POST("/reservations/:reservationId/email-confirmation", handlers.CreateTicketRelease)

	fmt.Println("Server started on http://localhost:" + os.Getenv("PORT"))
	http.ListenAndServe(":" + os.Getenv("PORT"), router)
}