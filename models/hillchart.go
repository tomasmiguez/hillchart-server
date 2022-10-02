package models

import "time"

type Hillchart struct {
	ID        string      `json:"id" gorm:"primaryKey;default:gen_random_uuid();"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Scopes    []Scope   `json:"scopes"`
	Frames    []Frame   `json:"frames"`
}
