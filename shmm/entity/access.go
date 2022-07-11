package entity

import (
	"time"

	"gorm.io/gorm"
)

type Access struct {
	gorm.Model

	ModelName  string
	ModelTag   string
	AccessedAt time.Time `gorm:"autoCreateTime"`

	// Many-to-One
	SharedModel SharedModel `gorm:"ForeignKey:ModelName,ModelTag;References:Name,Tag"`
}
