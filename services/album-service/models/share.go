package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Share struct {
	gorm.Model
	TokenID     uuid.UUID `gorm:"type:uuid;"`
	OwnerID     uint64
	Album       Album `gorm:"foreignkey:ID"`
	AccessedAt  time.Time
	ExpiresAt   time.Time
	Status      ShareStatus
	Permissions SharePermissions
}

type ShareStatus string

const (
	ShareStatusActive     ShareStatus = "ACTIVE"
	ShareStatusTerminated ShareStatus = "TERMINATED"
	ShareStatusExpired    ShareStatus = "EXPIRED"
)

func (s *Share) CheckStatus() {
	if s.ExpiresAt.After(time.Now()) {
		s.Status = ShareStatusExpired
	}
}

type SharePermissions string

const (
	SharePermissionsRead     SharePermissions = "READ"
	SharePermissionsWrite    SharePermissions = "WRITE"
	SharePermissionReadWrite SharePermissions = "READ_WRITE"
)
