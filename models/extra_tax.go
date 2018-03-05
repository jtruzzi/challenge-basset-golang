package models

type ExtraTax struct {
	Total float32 `json:"total,omitempty"`
	Type  string  `json:"type,omitempty"`
}
