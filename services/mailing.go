package services

import (
	"../models"
	"github.com/mostafah/mandrill"
	"log"
	"os"
)

func SendEmailConfirmation(reservation models.Reservation, product models.Product, flightReservation models.FlightReservation) ([]*mandrill.SendResult, error) {
	mandrill.Key = os.Getenv("BASSET_MANDRILL_API_KEY")
	// you can test your API key with Ping
	pingErr := mandrill.Ping()
	if pingErr != nil { log.Panic(pingErr) }

	data := map[string]string{
		"name": product.Passengers[0].FirstName,
		"reservation_code": flightReservation.PNR,
	}
	message := mandrill.NewMessageTo("julio.truzzi@gmail.com","Julio Truzzi")

	attach := &mandrill.Attachment{
		"application/pdf",
		"vuelo.pdf",
		GenerateFlightConfirmationPDF(reservation, product, flightReservation),
	}
	message.Attachments = []*mandrill.Attachment{attach}
	response, err := message.SendTemplate("issued-ticket", data, false)

	log.Println("Id", response[0].Id)
	log.Println("Email", response[0].Email)
	log.Println("Status", response[0].Status)
	log.Println("RejectionReason", response[0].RejectionReason)

	if err != nil { log.Panic(err) }

	return response, err
}

