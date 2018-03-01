package services

import (
	"../models"
	"github.com/mostafah/mandrill"
	"log"
	"os"
)

func SendEmailConfirmation(reservation models.Reservation, products []models.Product) ([]*mandrill.SendResult, error) {

	// Loop over flightReservations and Generate attachments contents as strings (base64) (Maybe Upload PDF to S3)
	var attachments []*mandrill.Attachment
	for _, product := range products {
		pdfContent := GenerateConfirmationPDF(reservation, product)
		attachment := &mandrill.Attachment{
			Mime:    "application/pdf",
			Name:    product.Type + ".pdf",
			Content: pdfContent,
		}
		attachments = append(attachments, attachment)
	}

	//
	// Send Email Confirmation (receiving attachments contents
	// Persist new ticket release in database


	mandrill.Key = os.Getenv("BASSET_MANDRILL_API_KEY")
	// you can test your API key with Ping
	pingErr := mandrill.Ping()
	if pingErr != nil { log.Panic(pingErr) }

	data := map[string]string{
		"name": products[0].Passengers[0].FirstName,
	}
	message := mandrill.NewMessageTo("julio.truzzi@gmail.com","Julio Truzzi")


	message.Attachments = attachments
	response, err := message.SendTemplate("issued-ticket", data, false)

	log.Println("Id", response[0].Id)
	log.Println("Email", response[0].Email)
	log.Println("Status", response[0].Status)
	log.Println("RejectionReason", response[0].RejectionReason)

	if err != nil { log.Panic(err) }

	return response, err
}
