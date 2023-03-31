package models

import "time"

type Car struct {
	ID        uint   `json:"id"`
	Model     string `json:"model"`
	Brand     string `json:"brand"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
