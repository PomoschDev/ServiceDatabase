package models

import (
	"DababaseService/pkg/database"
	"gorm.io/gorm"
)

func AllCards() ([]*Card, error) {
	db := database.GetDB()
	var cards []*Card

	err := db.Find(&cards).Error

	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (c *Card) Create() error {
	db := database.GetDB()

	return db.Create(&c).Error
}

func (c *Card) FindCardID() error {
	db := database.GetDB()

	return db.Where(&Card{ID: c.ID}).Find(&c).Error
}

func (c *Card) Delete() error {
	db := database.GetDB()

	return db.Delete(&c).Error
}

func (c *Card) DeleteByID() error {
	dbase := database.GetDB()
	result := dbase.Delete(&Card{}, c.ID)

	return result.Error
}

func (c *Card) Update() error {
	dbase := database.GetDB()
	result := dbase.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&c)
	dbase.Save(&c)
	return result.Error
}
