package models

type Leg struct {
	Origin           Airport   `json:"origin,omitempty"`
	Destination      Airport   `json:"destination,omitempty"`
	DepartureDate    string    `json:"departure_date,omitempty"`
	DepartureTime    string    `json:"departure_time,omitempty"`
	ArrivalDate      string    `json:"arrival_date,omitempty"`
	ArrivalTime      string    `json:"arrival_time,omitempty"`
	MarketingCarrier Carrier   `json:"marketing_carrier,omitempty"`
	OperatingCarrier Carrier   `json:"operating_carrier,omitempty"`
	Duration         string    `json:"duration,omitempty"`
	FlightNumber     string    `json:"flight_number,omitempty"`
	AircraftType     string    `json:"aircraft_type,omitempty"`
	TechnicalStop    string    `json:"technical_stop,omitempty"`
	CabinType        CabinType `json:"cabin_type,omitempty"`
	Status           string    `json:"status,omitempty"`
	RecordLocator    string    `json:"record_locator,omitempty"`
}
