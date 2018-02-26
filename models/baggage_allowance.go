package models // import "github.com/jtruzzi/basset-mailing-gateway/models"

type BaggageAllowance struct {
	Quantity int `json:"quantity,omitempty"`
	Weight int `json:"weight,omitempty"`
}
