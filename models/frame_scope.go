package models

import "time"

type FrameScope struct {
	ID        string      `json:"id" gorm:"primaryKey;default:gen_random_uuid();"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title     string    `json:"title"`
	Position  float32   `json:"position" gorm:"default:0"`
	FrameID   string      `json:"frame_id"`
	ScopeID   string      `json:"scope_id"`
}
