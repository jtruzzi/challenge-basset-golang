package models

type Fare struct {
	Currency       string     `json:"currency,omitempty"`
	ConversionRate float32    `json:"conversion_rate,omitempty"`
	Total          float32    `json:"total,omitempty"`
	TotalTax       float32    `json:"total_tax,omitempty"`
	BaseFare       float32    `json:"base_fare,omitempty"`
	ExtraTaxes     []ExtraTax `json:"extra_taxes,omitempty"`
}
