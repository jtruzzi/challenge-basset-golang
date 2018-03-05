package services

import (
	"errors"
	"log"
	"os"

	"../models"
	"github.com/mostafah/mandrill"
)

func SendEmailConfirmation(reservation models.Reservation, products []models.Product, resend bool) ([]*mandrill.SendResult, error) {
	mandrill.Key = os.Getenv("BASSET_MANDRILL_API_KEY")
	pingErr := mandrill.Ping()
	if pingErr != nil {
		log.Panic(pingErr)
	}

	var attachments []*mandrill.Attachment
	for _, product := range products {
		attachment := generateAttachments(product, reservation, resend)
		if (attachment != nil) {
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
			"name": products[0].Passengers[0].FirstName,
		}
		message.AddGlobalMergeVars(globalVars)
		message.MergeLanguage = "handlebars"

		return message.SendTemplate("issued-ticket-email", nil, false)
	}
	return nil, errors.New("no attachments to be sent")
}

func generateAttachments(product models.Product, reservation models.Reservation, resend bool) *mandrill.Attachment {
	ticketRelease, _ := models.GetTicketRelease(product.ItemId)
	if resend == true || ticketRelease.Released != true {
		var attachment models.Attachment

		if ticketRelease.S3Url != "" {
			attachment, _ = GetAttachmentFromS3(ticketRelease.S3Url)
		} else {
			attachment, _ = GenerateConfirmationPDF(reservation, product)
			ticketRelease.S3Url = SaveAttachmentToS3(attachment)
			ticketRelease.Released = true
			ticketRelease.Save()
		}

		return generateMandrillAttachment(attachment)
	}
	return nil
}

func generateMandrillAttachment(a models.Attachment) *mandrill.Attachment {
	return &mandrill.Attachment{
		Mime:    a.Mime,
		Name:    a.Name(),
		Content: a.Base64Content(),
	}
}
