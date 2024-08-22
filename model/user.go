package model

import (
	"time"

	"github.com/google/uuid"
)

// User struct
type User struct {
	ID         uuid.UUID `gorm:"primarykey" json:"id"`
	Active     bool      `json:"active" example:"true"`
	CreatedBy  string    `json:"created_by" binding:"required"  example:"ankit"`
	Email      string    `json:"email" binding:"required" gorm:"unique;not null" example:"ankit"`
	Password   string    `json:"password" binding:"required" gorm:"not null"  example:"password"`
	CreatedAt  time.Time `json:"created_at" `
	UpdatedBy  string    `json:"updated_by" `
	UpdatedAt  time.Time `json:"updated_at" `
	DeletedBy  string    `json:"deleted_by" `
	DeletedAt  time.Time `json:"deleted_at" `
	FirstName  string    `json:"first_name" binding:"required" example:"John"`
	MiddleName string    `json:"middle_name" binding:"required"  example:"Doe"`
	LastName   string    `json:"last_name" binding:"required" example:"Smith"`
	Lane       string    `json:"lane" example:"1234 Elm St"`
	Village    string    `json:"village" example:"Springfield"`
	City       string    `json:"city" binding:"required" example:"Metropolis"`
	District   string    `json:"district" binding:"required" example:"Central"`
	Pincode    int       `json:"pincode" binding:"required" example:"123456"`
	State      string    `json:"state" binding:"required" example:"NY"`
	Type       string    `json:"type" gorm:"not null"`
}

// UserSignIn struct
type UserSignIn struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
