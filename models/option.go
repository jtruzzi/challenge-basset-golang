package models

type Option struct {
	Id               string             `json:"id,omitempty"`
	DepartureDate    string             `json:"departure_date,omitempty"`
	DapartureTime    string             `json:"departure_time,omitempty"`
	ArrivalDate      string             `json:"arrival_date,omitempty"`
	ArrivalTime      string             `json:"arrival_time,omitempty"`
	Duration         string             `json:"duration,omitempty"`
	Legs             []Leg              `json:"legs,omitempty"`
	BaggageAllowance []BaggageAllowance `json:"baggage_allowance,omitempty"`
}
