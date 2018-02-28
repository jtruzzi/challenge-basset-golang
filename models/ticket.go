package models

type Ticket struct {
	Number string `json:"number,omitempty"`
	Passenger string `json:"passenger,omitempty"`
	EmissionDate string `json:"emission_date,omitempty"`
	PassengerType string `json:"passenger_type,omitempty"`
	Status string `json:"status,omitempty"`
}

func (ticket Ticket) Issued() bool {
	return ticket.Status == "ISSUED"
}
