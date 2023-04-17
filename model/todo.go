package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model           // Adds some metadata fields to the table
	ID         uuid.UUID `gorm:"type:uuid"` // Explicitly specify the type to be uuid
	Title      string
	SubTitle   string
	Text       string
}
