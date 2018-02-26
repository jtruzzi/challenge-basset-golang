package models // import "github.com/jtruzzi/basset-mailing-gateway/models"

type Segment struct {
	Origin Airport `json:"origin,omitempty"`
	Destination Airport `json:"destination,omitempty"`
	DepartureDate string `json:"departure_date,omitempty"`
	Options []Option `json:"options,omitempty"`
}
