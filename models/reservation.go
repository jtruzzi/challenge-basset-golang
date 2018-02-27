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
