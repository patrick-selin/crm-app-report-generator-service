package models

import (
	"time"
)

type ReportRequest struct {
	OrderIDs     []string `json:"order_ids" validate:"required,dive,uuid"`
	ReportType   string   `json:"report_type" validate:"required,oneof=csv pdf both"`
	IncludeItems bool     `json:"include_items"`
}


type Report struct {
	ReportID     string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"report_id"`
	FileURL      string    `json:"file_url"`
	Status       string    `json:"status"`
	ReportType   string    `json:"report_type"`
	IncludeItems bool      `json:"include_items"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
