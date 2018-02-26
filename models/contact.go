package models // import "github.com/jtruzzi/basset-mailing-gateway/models"

type Contact struct {
	Language string `json:"language,omitempty"`
	Email string `json:"email,omitempty"`
	Telephone Telephone `json:"telephone,omitempty"`
}
