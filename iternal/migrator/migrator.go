package migrator

import (
	"DababaseService/pkg/database/models"
	"errors"
	"gorm.io/gorm"
)

func Migrations(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}, &models.Card{}, &models.Company{}, &models.CardCompany{}, &models.Ward{},
		&models.Donations{}, &models.Session{})
	if err != nil {
		return errors.New("Migration error: " + err.Error())
	}

	return nil
}
