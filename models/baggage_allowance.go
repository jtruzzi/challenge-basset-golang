package models

type BaggageAllowance struct {
	Quantity int `json:"quantity,omitempty"`
	Weight   int `json:"weight,omitempty"`
}
