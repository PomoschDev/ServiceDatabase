package models

import (
	"DababaseService/pkg/database"
	"gorm.io/gorm"
)

func AllDonations() ([]*Donations, error) {
	db := database.GetDB()
	var donations []*Donations

	err := db.Find(&donations).Error

	if err != nil {
		return nil, err
	}

	return donations, nil
}

func (d *Donations) Create() error {
	db := database.GetDB()
	result := db.Create(&d)

	return result.Error
}

func FindDonationWards(id uint64) (*Ward, error) {
	db := database.GetDB()
	d := new(Donations)
	err := db.Preload("Ward").Where(&Donations{ID: id}).Find(&d).Error

	if err != nil {
		return nil, err
	}

	if d.ID == 0 {
		return &Ward{}, nil
	}

	return d.Ward, err
}

func (d *Donations) FindDonationID() error {
	db := database.GetDB()

	return db.Where(&Donations{ID: d.ID}).Find(&d).Error
}

func (d *Donations) Delete() error {
	db := database.GetDB()

	return db.Select("Ward").Delete(&d).Error
}

func (d *Donations) DeleteByID() error {
	dbase := database.GetDB()
	result := dbase.Select("Ward").Delete(&Donations{}, d.ID)

	return result.Error
}

func (d *Donations) Update() error {
	dbase := database.GetDB()
	result := dbase.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&d)
	dbase.Save(&d)
	return result.Error
}
