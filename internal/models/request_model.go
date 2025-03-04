package models

// ReportRequest represents the request body for creating a report
type ReportRequest struct {
	OrderIDs     []string `json:"order_ids" validate:"required,min=1,dive,uuid4"`
	ReportType   string   `json:"report_type" validate:"required,oneof=csv pdf both"`
	IncludeItems bool     `json:"include_items"`
}
