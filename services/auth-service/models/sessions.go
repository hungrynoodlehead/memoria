package models

import "gorm.io/gorm"

type SessionType string

const (
	Disabled   SessionType = "DISABLED"
	Active     SessionType = "ACTIVE"
	Terminated SessionType = "TERMINATED"
)

type Sessions struct {
	gorm.Model
	User             User
	UserID           uint
	FirstUserAgent   string
	CurrentUserAgent string
	FirstIP          string
	CurrentIP        string
	Status           SessionType  `gorm:"type:session_status"`
	TokenPairs       []TokenPairs `gorm:"foreignKey:ID"`
}
