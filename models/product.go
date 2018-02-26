package models // import "github.com/jtruzzi/basset-mailing-gateway/models"

type Product struct {
	ReservationId string `json:"reservation_id,omitempty"`
	Type string `json:"type,omitempty"`
	ItemId string `json:"item_id,omitempty"`
	Status string `json:"status,omitempty"`
	Passengers []Passenger `json:"passengers,omitempty"`
	Fare Fare `json:"fare,omitempty"`
	FlightReservation FlightReservation `json:"flight_reservation,omitempty"`
}