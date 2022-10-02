package models

import "time"

type Frame struct {
	ID        string      `json:"id" gorm:"primaryKey;default:gen_random_uuid();"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	HillchartID string         `json:"hillchart_id"`
	FrameScopes []FrameScope `json:"frame_scopes"`
}
