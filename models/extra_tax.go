package models // import "github.com/jtruzzi/basset-mailing-gateway/models"

type ExtraTax struct {
	Total float32 `json:"total,omitempty"`
	Type string `json:"type,omitempty"`
}
