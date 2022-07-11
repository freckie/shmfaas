package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/freckie/shmfaas/shmm/entity"
)

func InitializeDB(filename string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(filename), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if mig := db.Migrator(); !mig.HasTable(&entity.SharedModel{}) {
		err = mig.CreateTable(&entity.SharedModel{})
		if err != nil {
			return nil, err
		}
	}
	if mig := db.Migrator(); !mig.HasTable(&entity.Access{}) {
		err = mig.CreateTable(&entity.Access{})
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
