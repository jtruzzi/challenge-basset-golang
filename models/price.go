package models // import "github.com/jtruzzi/basset-mailing-gateway/models"

type Price struct {
	Adults Adults `json:"adults,omitempty"`
	charges float32 `json:"charges,omitempty"`
	Currency string `json:"currency,omitempty"`
	Fees float32 `json:"fees,omitempty"`
	Total float32 `json:"total,omitempty"`
}
