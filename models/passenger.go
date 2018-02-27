package models

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