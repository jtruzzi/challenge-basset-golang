package models

type FlightReservation struct {
	Id string `json:"id"`
	ClientId string `json:"client_id"`
	Cluster Cluster `json:"cluster"`
	Contact Contact `json:"contact"`
	FiscalIdentification FiscalIdentification `json:"fiscal_identification"`
	Passengers []Passenger `json:"passengers"`
	PNR string `json:"pnr"`
	Site string `json:"site"`
	Status string `json:"status"`
	Tickets []Ticket `json:"tickets"`
}

type Ticket struct {
	Number string `json:"number"`
	Passenger string `json:"passenger"`
	EmissionDate string `json:"emission_date"`
	PassengerType string `json:"passenger_type"`
	Status string `json:"status"`
}

type Cluster struct {
	Segments []Segment `json:"segments"`
	Price Price `json:"price"`
	ValidatingCarrier Carrier `json:"validating_carrier"`
}


type Price struct {
	Adults Adults `json:"adults"`
	charges float32 `json:"charges"`
	Currency string `json:"currency"`
	Fees float32 `json:"fees"`
	Total float32 `json:"total"`
}

type Adults struct {
	Fare int `json:"fare"`
	Quantity int `json:"quantity"`
}

type Segment struct {
	Origin Airport `json:"origin"`
	Destination Airport `json:"destination"`
	DepartureDate string `json:"departure_date"`
	Options []Option `json:"options"`
}

type Option struct {
	Id string `json:"id"`
	DepartureDate string `json:"departure_date"`
	DapartureTime string `json:"departure_time"`
	ArrivalDate string `json:"arrival_date"`
	ArrivalTime string `json:"arrival_time"`
	Duration string `json:"duration"`
	Legs []Leg `json:"legs"`
	BaggageAllowance []BaggageAllowance `json:"baggage_allowance"`
}

type BaggageAllowance struct {
	Quantity int `json:"quantity"`
	Weight int `json:"weight"`
}

type Leg struct {
	Origin Airport `json:"origin"`
	Destination Airport `json:"destination"`
	DepartureDate string `json:"departure_date"`
	DepartureTime string `json:"departure_time"`
	ArrivalDate string `json:"arrival_date"`
	ArrivalTime string `json:"arrival_time"`
	MarketingCarrier Carrier `json:"marketing_carrier"`
	OperatingCarrier Carrier `json:"operating_carrier"`
	Duration string `json:"duration"`
	FlightNumber string `json:"flight_number"`
	AircraftType string `json:"aircraft_type"`
	TechnicalStop string `json:"technical_stop"`
	CabinType CabinType `json:"cabin_type"`
	Status string `json:"status"`
	RecordLocator string `json:"record_locator"`
}

type CabinType struct {
	code string `json:"code"`
	name string `json:"name"`
}

type Airport struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Carrier struct {
	Code string `json:"code"`
	Name string `json:"name"`
}