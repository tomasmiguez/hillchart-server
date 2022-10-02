package models

import "time"

type Scope struct {
	ID        string      `json:"id" gorm:"primaryKey;default:gen_random_uuid();"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Color       string    `json:"color"`
	HillchartID string      `json:"hillchart_id"`
}
