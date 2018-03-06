package services

import (
	"errors"
	"log"

	"../models"
	"github.com/mostafah/mandrill"
)

// SendEmailConfirmation: Sends email configuration with attachments for all issued products
func SendEmailConfirmation(reservation models.Reservation, resend bool, client models.Client) ([]*mandrill.SendResult, error) {
	mandrill.Key = client.MandrillApiKey
	pingErr := mandrill.Ping()
	if pingErr != nil {
		log.Panic(pingErr)
	}

	var attachments []*mandrill.Attachment
	for _, product := range reservation.Products {
		attachment := generateAttachment(product, reservation, resend, client)
		if attachment != nil {
			attachments = append(attachments, attachment)
		}
	}

	// TODO: Uncomment line to use actual reservation email
	//email := reservation.Contact.Email
	email := "julio.truzzi@gmail.com"
	if len(attachments) > 0 {
		message := mandrill.NewMessageTo(email, "")

		message.Attachments = attachments
		globalVars := map[string]interface{}{
			"name":             reservation.Products[0].Passengers[0].FirstName,
			"client_name":      client.Name,
			"reservation_code": reservation.Products[0].FlightReservation.PNR,
		}
		message.AddGlobalMergeVars(globalVars)
		message.MergeLanguage = "handlebars"

		return message.SendTemplate("issued-ticket-email", nil, false)
	}
	return nil, errors.New("no attachments to be sent")
}

// generateAttachments: Generates mandrill attachment based on a given product
func generateAttachment(product models.Product, reservation models.Reservation, resend bool, client models.Client) *mandrill.Attachment {
	ticketRelease, _ := models.GetTicketRelease(product.ItemId)
	if resend == true || ticketRelease.Released != true {
		var attachment models.Attachment

		if ticketRelease.S3Url != "" {
			attachment, _ = GetAttachmentFromS3(ticketRelease.S3Url)
		} else {
			attachment, _ = GenerateConfirmationPDF(reservation, product, client)
			ticketRelease.S3Url = SaveAttachmentToS3(attachment)
			ticketRelease.Released = true
			ticketRelease.Save()
		}

		return generateMandrillAttachment(attachment)
	}
	return nil
}

// generateMandrillAttachment: Transform a basset attachment into a mandrill attachment
func generateMandrillAttachment(a models.Attachment) *mandrill.Attachment {
	return &mandrill.Attachment{
		Mime:    a.Mime,
		Name:    a.Name(),
		Content: a.Base64Content(),
	}
}
