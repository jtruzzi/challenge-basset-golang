package models

type Telephone struct {
	Type        string `json:"type,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	AreaCode    string `json:"area_code,omitempty"`
	Number      string `json:"number,omitempty"`
}
