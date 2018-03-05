package models

type Price struct {
	Adults   Adults  `json:"adults,omitempty"`
	Charges  float32 `json:"charges,omitempty"`
	Currency string  `json:"currency,omitempty"`
	Feew     float32 `json:"fees,omitempty"`
	Total    float32 `json:"total,omitempty"`
}
