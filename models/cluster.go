package models

type Cluster struct {
	Segments []Segment `json:"segments,omitempty"`
	Price Price `json:"price,omitempty"`
	ValidatingCarrier Carrier `json:"validating_carrier,omitempty"`
}
