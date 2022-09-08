package models

import "time"

type Frame struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	HillchartID uint `json:"hillchart_id"`
	Scopes []Scope
}
