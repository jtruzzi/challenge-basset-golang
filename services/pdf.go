package services // import "github.com/jtruzzi/basset-mailing-gateway/services"

import (
	"github.com/jtruzzi/basset-mailing-gateway/models"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"bytes"
	"strings"
	"log"
	"html/template"
	"encoding/base64"
)

func GenerateFlightConfirmationPDF(reservation models.Reservation, product models.Product, flightReservation models.FlightReservation) string {
	generator, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	generator.Dpi.Set(600)
	generator.NoCollate.Set(false)
	generator.PageSize.Set(wkhtmltopdf.PageSizeA4)
	generator.MarginBottom.Set(40)

	tpl, err := template.ParseFiles("templates/flights/issued-ticket.html")
	if err != nil {
		log.Fatalln(err)
	}
	buf := new(bytes.Buffer)

	err = tpl.Execute(buf, nil)
	if err != nil {
		log.Fatalln(err)
	}

	generator.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(buf.String())))

	err = generator.Create()
	if err != nil {
		log.Fatal(err)
	}

	return base64.StdEncoding.EncodeToString(generator.Bytes())
}