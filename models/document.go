package models // import "github.com/jtruzzi/basset-mailing-gateway/models"

type Document struct {
	Type string `json:"type,omitempty"`
	Number string `json:"number,omitempty"`
}