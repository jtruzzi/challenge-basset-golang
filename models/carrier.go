package models // import "github.com/jtruzzi/basset-mailing-gateway/models"

type Carrier struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}