package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"../models"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type MandrillTemplateResponse struct {
	Slug        string `json:"type"`
	Name        string `json:"firstname"`
	PublishCode string `json:"publish_code"`
}

func GenerateConfirmationPDF(reservation models.Reservation, product models.Product) (models.Attachment, error) {
	pdfGenerator, _ := wkhtmltopdf.NewPDFGenerator()
	pdfGenerator.Dpi.Set(600)
	pdfGenerator.NoCollate.Set(false)
	pdfGenerator.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfGenerator.MarginBottom.Set(40)

	issuedTicketHeader, _ := template.ParseFiles(fmt.Sprintf("templates/%s/issued-ticket.gohtml", strings.ToLower(product.Type)))

	buffer := new(bytes.Buffer)

	err := issuedTicketHeader.Execute(buffer, map[string]interface{}{
		"Header":      fetchTemplate("issued-ticket-pdf-header"),
		"Footer":      fetchTemplate("issued-ticket-pdf-footer"),
		"Product":     product,
		"Reservation": reservation,
	})

	if err != nil {
		log.Println(err)
		return models.Attachment{}, err
	}

	pdfGenerator.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(buffer.String())))

	err = pdfGenerator.Create()
	if err != nil {
		log.Fatal(err)
	}

	return models.Attachment{
		Mime:    "application/pdf",
		Path:    product.FlightReservation.PNR + ".pdf",
		Content: pdfGenerator.Bytes(),
	}, nil
}

func fetchTemplate(name string) template.HTML {
	client := &http.Client{}
	var body = []byte(`{"key":"` + os.Getenv("BASSET_MANDRILL_API_KEY") + `", "name": "` + name + `"}`)

	request, err := http.NewRequest("POST", "https://mandrillapp.com/api/1.0/templates/info.json", bytes.NewBuffer(body))
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The http request failed with error %s/n", err)
	}

	data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	var templateResponse MandrillTemplateResponse
	errMarshall := json.Unmarshal([]byte(data), &templateResponse)
	if errMarshall != nil {
		return ""
	}
	return template.HTML(templateResponse.PublishCode)
}
