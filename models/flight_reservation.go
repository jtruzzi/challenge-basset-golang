package models

type FlightReservation struct {
	Id string `json:"id,omitempty"`
	ClientId string `json:"client_id,omitempty"`
	Cluster Cluster `json:"cluster,omitempty"`
	Contact Contact `json:"contact,omitempty"`
	FiscalIdentification FiscalIdentification `json:"fiscal_identification,omitempty"`
	Passengers []Passenger `json:"passengers,omitempty"`
	PNR string `json:"pnr,omitempty"`
	Site string `json:"site,omitempty"`
	Status string `json:"status,omitempty"`
	Tickets []Ticket `json:"tickets,omitempty"`
}

type Ticket struct {
	Number string `json:"number,omitempty"`
	Passenger string `json:"passenger,omitempty"`
	EmissionDate string `json:"emission_date,omitempty"`
	PassengerType string `json:"passenger_type,omitempty"`
	Status string `json:"status,omitempty"`
}

type Cluster struct {
	Segments []Segment `json:"segments,omitempty"`
	Price Price `json:"price,omitempty"`
	ValidatingCarrier Carrier `json:"validating_carrier,omitempty"`
}


type Price struct {
	Adults Adults `json:"adults,omitempty"`
	charges float32 `json:"charges,omitempty"`
	Currency string `json:"currency,omitempty"`
	Fees float32 `json:"fees,omitempty"`
	Total float32 `json:"total,omitempty"`
}

type Adults struct {
	Fare int `json:"fare,omitempty"`
	Quantity int `json:"quantity,omitempty"`
}

type Segment struct {
	Origin Airport `json:"origin,omitempty"`
	Destination Airport `json:"destination,omitempty"`
	DepartureDate string `json:"departure_date,omitempty"`
	Options []Option `json:"options,omitempty"`
}

type Option struct {
	Id string `json:"id,omitempty"`
	DepartureDate string `json:"departure_date,omitempty"`
	DapartureTime string `json:"departure_time,omitempty"`
	ArrivalDate string `json:"arrival_date,omitempty"`
	ArrivalTime string `json:"arrival_time,omitempty"`
	Duration string `json:"duration,omitempty"`
	Legs []Leg `json:"legs,omitempty"`
	BaggageAllowance []BaggageAllowance `json:"baggage_allowance,omitempty"`
}

type BaggageAllowance struct {
	Quantity int `json:"quantity,omitempty"`
	Weight int `json:"weight,omitempty"`
}

type Leg struct {
	Origin Airport `json:"origin,omitempty"`
	Destination Airport `json:"destination,omitempty"`
	DepartureDate string `json:"departure_date,omitempty"`
	DepartureTime string `json:"departure_time,omitempty"`
	ArrivalDate string `json:"arrival_date,omitempty"`
	ArrivalTime string `json:"arrival_time,omitempty"`
	MarketingCarrier Carrier `json:"marketing_carrier,omitempty"`
	OperatingCarrier Carrier `json:"operating_carrier,omitempty"`
	Duration string `json:"duration,omitempty"`
	FlightNumber string `json:"flight_number,omitempty"`
	AircraftType string `json:"aircraft_type,omitempty"`
	TechnicalStop string `json:"technical_stop,omitempty"`
	CabinType CabinType `json:"cabin_type,omitempty"`
	Status string `json:"status,omitempty"`
	RecordLocator string `json:"record_locator,omitempty"`
}

type CabinType struct {
	code string `json:"code,omitempty"`
	name string `json:"name,omitempty"`
}

type Airport struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type Carrier struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}