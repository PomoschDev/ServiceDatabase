package models

import (
	"DababaseService/pkg/database"
	"gorm.io/gorm"
)

func AllCardsCompanies() ([]*CardCompany, error) {
	db := database.GetDB()
	var cardsCompanies []*CardCompany

	err := db.Find(&cardsCompanies).Error

	if err != nil {
		return nil, err
	}

	return cardsCompanies, nil
}

func (cc *CardCompany) Create() error {
	db := database.GetDB()

	return db.Create(&cc).Error
}

func (cc *CardCompany) FindCardCompanyID() error {
	db := database.GetDB()

	return db.Where(&CardCompany{ID: cc.ID}).Find(&cc).Error
}

func (cc *CardCompany) Delete() error {
	db := database.GetDB()

	return db.Delete(&cc).Error
}

func (cc *CardCompany) DeleteByID() error {
	dbase := database.GetDB()
	result := dbase.Delete(&CardCompany{}, cc.ID)

	return result.Error
}

func (cc *CardCompany) Update() error {
	dbase := database.GetDB()
	result := dbase.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&cc)
	dbase.Save(&cc)
	return result.Error
}
