package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Model struct {
	ID        uuid.UUID  `sql:"type:uuid;default:uuid_generate_v4()" json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
}

type Entry struct {
	Model    `json:",inline"`
	Property bool `json:"property"`
}
