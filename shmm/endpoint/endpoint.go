package endpoint

import (
	"gorm.io/gorm"
)

type Endpoint struct {
	DB *gorm.DB
}
