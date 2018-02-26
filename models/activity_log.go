package models // import "github.com/jtruzzi/basset-mailing-gateway/models"

type ActivityLog struct {
	ProductId string `json:"product_id,omitempty"`
	Action string `json:"action,omitempty"`
	Product string `json:"product,omitempty"`
	PreviousStatus string `json:"previous_status,omitempty"`
	NewStatus string `json:"new_status,omitempty"`
	UserId string `json:"user_id,omitempty"`
	DateTime string `json:"date_time,omitempty"`
}
