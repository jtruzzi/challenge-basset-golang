package services

import (
	"../models"
	"github.com/keighl/mandrill"
	"log"
	"os"
)

var client *mandrill.Client = mandrill.ClientWithKey(os.Getenv("BASSET_MANDRILL_API_KEY"))

func SendEmailConfirmation(
		reservation models.Reservation,
		product models.Product,
		flightReservation models.FlightReservation,
	) (
		[]*mandrill.Response,
		error,
	) {
	message := &mandrill.Message{}
	message.AddRecipient("julio.truzzi@basset.la", "Bob Johnson", "to")
	message.FromEmail = "reservas@tuhotelhoy.com"
	message.FromName = "Tuhotelhoy.com"
	message.Subject = "Confirmación de reserva"

	templateContent := map[string]string{
		"header": "Confirmación de reserva",
	}
	responses, err := client.MessagesSendTemplate(message, "issued-ticket", templateContent)
	log.Print(responses)
	log.Print(err)
	return responses, err
}