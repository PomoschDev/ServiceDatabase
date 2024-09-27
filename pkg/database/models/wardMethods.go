package models

import (
	"DababaseService/pkg/database"
	"gorm.io/gorm"
)

func AllWards() ([]*Ward, error) {
	db := database.GetDB()
	var wards []*Ward

	err := db.Find(&wards).Error

	if err != nil {
		return nil, err
	}

	return wards, nil
}

func (w *Ward) Create() error {
	db := database.GetDB()

	return db.Create(&w).Error
}

func (w *Ward) FindWardID() error {
	db := database.GetDB()

	return db.Where(&Ward{ID: w.ID}).Find(&w).Error
}

func (w *Ward) Delete() error {
	db := database.GetDB()

	return db.Delete(&w).Error
}

func (w *Ward) DeleteByID() error {
	dbase := database.GetDB()
	result := dbase.Delete(&Ward{}, w.ID)

	return result.Error
}

func (w *Ward) Update() error {
	dbase := database.GetDB()
	result := dbase.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&w)
	dbase.Save(&w)
	return result.Error
}
