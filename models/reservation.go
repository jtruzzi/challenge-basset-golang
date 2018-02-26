package models

type Reservation struct {
	Id string `json:"id,omitempty"`
	Products []Product `json:"products,omitempty"`
	ClientId string `json:"client_id,omitempty"`
	Site string `json:"site,omitempty"`
	Channel string `json:"channel,omitempty"`
	Contact Contact `json:"contact,omitempty"`
	FiscalIdentification FiscalIdentification `json:"fiscal_identification,omitempty"`
	ActivityLogs []ActivityLog `json:"activity_logs,omitempty"`
}

type Product struct {
	ReservationId string `json:"reservation_id,omitempty"`
	Type string `json:"type,omitempty"`
	ItemId string `json:"item_id,omitempty"`
	Status string `json:"status,omitempty"`
	Passengers []Passenger `json:"passengers,omitempty"`
	Fare Fare `json:"fare,omitempty"`
	FlightReservation FlightReservation `json:"flight_reservation,omitempty"`
}

type Passenger struct {
	Type string `json:"type,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName string `json:"lastname,omitempty"`
	Gender string `json:"gender,omitempty"`
	Birth string `json:"birth,omitempty"`
	Document Document `json:"document,omitempty"`
	Nationality string `json:"nationality,omitempty"`
}

func (p Passenger) FullName() string {
	return p.FirstName + " " + p.LastName
}

type Document struct {
	Type string `json:"type,omitempty"`
	Number string `json:"number,omitempty"`
}

type Fare struct {
	Currency string `json:"currency,omitempty"`
	ConversionRate float32 `json:"conversion_rate,omitempty"`
	Total float32 `json:"total,omitempty"`
	TotalTax float32 `json:"total_tax,omitempty"`
	BaseFare float32 `json:"base_fare,omitempty"`
	ExtraTaxes []ExtraTax `json:"extra_taxes,omitempty"`
}

type ExtraTax struct {
	Total float32 `json:"total,omitempty"`
	Type string `json:"type,omitempty"`
}

type Contact struct {
	Language string `json:"language,omitempty"`
	Email string `json:"email,omitempty"`
	Telephone Telephone `json:"telephone,omitempty"`
}

type Telephone struct {
	Type string `json:"type,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	AreaCode string `json:"area_code,omitempty"`
	Number string `json:"number,omitempty"`
}

type FiscalIdentification  struct {
	Type string `json:"type,omitempty"`
	Number string `json:"number,omitempty"`
}

type ActivityLog struct {
	ProductId string `json:"product_id,omitempty"`
	Action string `json:"action,omitempty"`
	Product string `json:"product,omitempty"`
	PreviousStatus string `json:"previous_status,omitempty"`
	NewStatus string `json:"new_status,omitempty"`
	UserId string `json:"user_id,omitempty"`
	DateTime string `json:"date_time,omitempty"`
}