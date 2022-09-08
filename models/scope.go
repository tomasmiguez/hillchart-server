package models

import "time"

type Scope struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title string `json:"title"`
	Position float32 `json:"position" gorm:"default:0"`
	FrameID uint `json:"frame_id"`
}
