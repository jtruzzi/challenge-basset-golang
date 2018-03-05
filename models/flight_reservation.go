package models

type FlightReservation struct {
	Id                   string               `json:"id,omitempty"`
	ClientId             string               `json:"client_id,omitempty"`
	Cluster              Cluster              `json:"cluster,omitempty"`
	Contact              Contact              `json:"contact,omitempty"`
	FiscalIdentification FiscalIdentification `json:"fiscal_identification,omitempty"`
	Passengers           []Passenger          `json:"passengers,omitempty"`
	PNR                  string               `json:"pnr,omitempty"`
	Site                 string               `json:"site,omitempty"`
	Status               string               `json:"status,omitempty"`
	Tickets              []Ticket             `json:"tickets,omitempty"`
}

func (flightReservation FlightReservation) HasIssuedTicket() bool {
	for _, ticket := range flightReservation.Tickets {
		if ticket.Issued() {
			return true
		}
	}
	// TODO: Uncomment for production
	return true
}
