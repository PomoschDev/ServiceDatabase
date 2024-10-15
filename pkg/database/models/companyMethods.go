package models

import (
	"DababaseService/pkg/database"
	"errors"
	"gorm.io/gorm"
)

func AllCompanies() ([]*Company, error) {
	db := database.GetDB()
	var companies []*Company

	err := db.Find(&companies).Error

	if err != nil {
		return nil, err
	}

	if companies == nil {
		return []*Company{}, nil
	}

	return companies, nil
}

func (c *Company) Create() error {
	db := database.GetDB()

	return db.Create(&c).Error
}

func (c *Company) FindCompanyID() error {
	db := database.GetDB()

	return db.Where(&Company{ID: c.ID}).Find(&c).Error
}

func (c *Company) FindCompanyPhone() error {
	db := database.GetDB()

	return db.Where(&Company{Phone: c.Phone}).Find(&c).Error
}

func FindCompanyCard(id uint64) (*CardCompany, error) {
	db := database.GetDB()
	company := new(Company)
	err := db.Preload("Card").Where(&Company{ID: id}).Find(&company).Error

	if err != nil {
		return nil, err
	}

	if company.ID == 0 {
		return nil, errors.New("Компания не найдена")
	}

	if company.Card == nil {
		return nil, errors.New("Информация о банковской карте отсутствует")
	}

	return company.Card, nil
}

func (c *Company) Delete() error {
	db := database.GetDB()

	return db.Select("Card").Delete(&c).Error
}

func (c *Company) DeleteByID() error {
	dbase := database.GetDB()
	result := dbase.Select("Card").Delete(&Company{}, c.ID)

	return result.Error
}

func (c *Company) Update() error {
	dbase := database.GetDB()
	result := dbase.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&c)
	dbase.Save(&c)
	return result.Error
}
