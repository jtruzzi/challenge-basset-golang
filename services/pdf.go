package services

import (
	"../models"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"bytes"
	"strings"
	"log"
	"html/template"
	"encoding/base64"
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)



type MandrillTemplateResponse struct {
	Slug string `json:"type"`
	Name string `json:"firstname"`
	PublishCode string `json:"publish_code"`
}

func GenerateConfirmationPDF(reservation models.Reservation, product models.Product) string {
	generator, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	generator.Dpi.Set(600)
	generator.NoCollate.Set(false)
	generator.PageSize.Set(wkhtmltopdf.PageSizeA4)
	generator.MarginBottom.Set(40)

	localTemplate, err := template.ParseFiles("templates/" + strings.ToLower(product.Type) + "/issued-ticket.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	localTemplateBuffer := new(bytes.Buffer)

	err = localTemplate.Execute(localTemplateBuffer, map[string]interface{} {
		"Header": fetchTemplate("issued-ticket-pdf-header"),
		"Footer": fetchTemplate("issued-ticket-pdf-footer"),
		"Product": product,
		"Reservation": reservation,
	})

	if err != nil { log.Println(err) }

	generator.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(localTemplateBuffer.String())))

	err = generator.Create()
	if err != nil {
		log.Fatal(err)
	}

	return base64.StdEncoding.EncodeToString(generator.Bytes())
}


func fetchTemplate(name string) template.HTML {
	client := &http.Client{}
	var body = []byte(`{"key":"`+ os.Getenv("BASSET_MANDRILL_API_KEY") +`", "name": "`+name+`"}`)
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
