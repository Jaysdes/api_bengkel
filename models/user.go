package models

import "time"

type User struct {
	ID              uint       `gorm:"primaryKey" json:"id"`
	Name            string     `gorm:"size:100" json:"name"`
	Email           string     `gorm:"size:100;unique" json:"email"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	Password        string     `gorm:"size:255" json:"-"`
	Role            string     `gorm:"type:enum('mekanik','admin','keuangan','gudang','customer');default:'customer'" json:"role"`
	RememberToken   string     `gorm:"size:100" json:"remember_token"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`

	Status string `gorm:"-" json:"status"`
}
