package model

import (
	"encoding/json"
	"time"
)

type Network struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Type 	  string 	`json:"type" gorm:"not null"`
	Nationality string 	`json:"nationality"`
	Ethnicity string 	`json:"ethnicity"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	UserId    uint      `json:"user_id" gorm:"not null"`
	Connections   json.RawMessage `json:"connections" gorm:"type:json"` // JSON 필드로 connections 저장

}

type NetworkResponse struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Type 	  string 	`json:"type" gorm:"not null"`
	Nationality string 	`json:"nationality"`
	Ethnicity string 	`json:"ethnicity"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Connections   json.RawMessage `json:"connections"` // JSON 필드로 connections 저장
}
