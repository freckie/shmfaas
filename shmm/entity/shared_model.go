package entity

import (
	"gorm.io/gorm"
)

type SharedModel struct {
	gorm.Model `json:"-"`

	Name     string `json:"name" gorm:"primaryKey"`
	Tag      string `json:"tag" gorm:"primaryKey"`
	Shmname  string `json:"shmname" gorm:"unique"`
	Shmsize  int64  `json:"shmsize"`
	Metadata []byte `json:"metadata"`
}
