package models

import "time"

type Movie struct {
	ID        string `gorm:"primarykey"`
	Title     string
	Genre     string
	CreatedAt time.Time `gorm:"autoCreateTime:false"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:false"`
}

type Movies struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"Title"`
	Genre string `json:"genre"`
}
