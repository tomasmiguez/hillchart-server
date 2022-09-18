package models

import "time"

type Scope struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Color       string    `json:"color"`
	HillchartID uint      `json:"hillchart_id"`
}
