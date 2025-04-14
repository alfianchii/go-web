package models

import "time"

type Role struct {
	ID int `json:"role_id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy int `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *int `json:"updated_by"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *int `json:"deleted_by"`
}