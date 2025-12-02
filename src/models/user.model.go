package models

import "time"

type UserModels struct {
	ID        int       `gorm:"primaryKey;increment"`
	Email     string    `gorm:"unique;not null;"`
	Username  string    `gorm:"unique;not null;"`
	Password  string    `gorm:"not null;size:100"`
	IsActive  bool      `gorm:"not null" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (UserModels) TableName() string {
	return "mst_users"
}
