package services

import (
	"../models"
	"github.com/mostafah/mandrill"
	"log"
	"os"
)

// TODO: Receive mail provider as a parameters as interface
func SendEmailConfirmation(reservation models.Reservation, products []models.Product) ([]*mandrill.SendResult, error) {
	mandrill.Key = os.Getenv("BASSET_MANDRILL_API_KEY")
	pingErr := mandrill.Ping()
	if pingErr != nil { log.Panic(pingErr) }

	message := mandrill.NewMessageTo("julio.truzzi@gmail.com","Julio Truzzi")

	message.Attachments = generateAttachments(products, reservation)

	message.AddGlobalMergeVars(map[string]interface{} {
		"name": products[0].Passengers[0].FirstName,
	})
	message.MergeLanguage = "handlebars"

	response, err := message.SendTemplate("issued-ticket-email", nil, false)

	if err != nil { log.Panic(err) }

	return response, err
}

func generateAttachments(products []models.Product, reservation models.Reservation) []*mandrill.Attachment {
	var attachments []*mandrill.Attachment
	for _, product := range products {
		// TODO: Persist new ticket release in database
		ticketRelease, _ := models.GetTicketRelease(product.ItemId)
		if ticketRelease.Released != true {
			pdfContent := GenerateConfirmationPDF(reservation, product)

			attachment := &mandrill.Attachment{
				Mime:    "application/pdf",
				Name:    product.Type + ".pdf",
				Content: pdfContent,
			}
			attachments = append(attachments, attachment)

			s3Url := saveAttachmentToS3(pdfContent)

			ticketRelease, _ = models.CreateTicketRelease(product.ItemId, true, s3Url)
		}
	}

	return attachments
}

func saveAttachmentToS3(pdfContent string) string {
	// TODO: Upload to S3
	return "dummy_s3_url"
}