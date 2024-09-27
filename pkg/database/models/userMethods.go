package models

import (
	"DababaseService/pkg/database"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const (
	Individual = iota
	Legal
)

func AllUsers() ([]*User, error) {
	db := database.GetDB()
	var users []*User

	err := db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) IsRole(role string) (bool, error) {
	if err := u.FindUserID(); err != nil {
		return false, err
	}

	if u.ID == 0 {
		return false, errors.New(fmt.Sprintf("Пользователь не найден"))
	}

	if u.Role == role {
		return true, nil
	}

	return false, nil
}

func (u *User) Create() error {
	db := database.GetDB()
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)

	return db.Create(&u).Error
}

func ComparePassword(phone string, password string) (bool, error) {
	u := &User{Phone: phone}
	if err := u.FindUserPhone(); err != nil {
		return false, err
	}

	if u.ID == 0 {
		return false, errors.New(fmt.Sprintf("Пользователь не найден"))
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}

	if string(hash) == u.Password {
		return true, nil
	}

	return false, nil
}

func UserIsExists(phone string) (bool, error) {
	u := &User{Phone: phone}
	if err := u.FindUserPhone(); err != nil {
		return false, err
	}

	if u.ID == 0 {
		return false, nil
	}

	return true, nil
}

func (u *User) FindUserID() error {
	db := database.GetDB()

	return db.Where(&User{ID: u.ID}).Find(&u).Error
}

func (u *User) FindUserEmail() error {
	db := database.GetDB()

	return db.Where(&User{Email: u.Email}).Find(&u).Error
}

func (u *User) FindUserPhone() error {
	db := database.GetDB()

	return db.Where(&User{Phone: u.Phone}).Find(&u).Error
}

func (u *User) ChangeUserType(newType uint64) error {
	db := database.GetDB()
	_ = db

	if err := u.FindUserID(); err != nil {
		return err
	}

	if u.ID == 0 {
		return errors.New("Пользователь не найден")
	}

	return db.UpdateColumn("type", newType).Error
}

func FindUserCompany(id uint64) (*Company, error) {
	db := database.GetDB()
	user := new(User)
	err := db.Preload("Company.Card").Where(&User{ID: id}).Find(&user).Error

	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, errors.New("Пользователь не найден")
	}

	if user.Type == Individual && user.Company == nil {
		return nil, errors.New("У юридических лиц отсутствуют компании")
	}

	if user.Type == Legal && user.Company == nil {
		return nil, errors.New("Компания не указана")
	}

	if user.Type == Legal && user.Company != nil {
		return user.Company, nil
	}

	return nil, nil
}

func FindUserDonations(id uint64) ([]*Donations, error) {
	db := database.GetDB()
	user := new(User)
	err := db.Preload("Donations").Where(&User{ID: id}).Find(&user).Error

	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, errors.New("Пользователь не найден")
	}

	if user.Donations == nil {
		return []*Donations{}, nil
	}

	if len(user.Donations) == 0 {
		return []*Donations{}, nil
	}

	return user.Donations, nil
}

func FindUserCard(id uint64) ([]*Card, error) {
	db := database.GetDB()
	user := new(User)
	err := db.Preload("Card").Where(&User{ID: id}).Find(&user).Error

	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, errors.New("Пользователь не найден")
	}

	if user.Card == nil {
		return nil, errors.New("Информация о банковской карте отсутствует")
	}

	if len(user.Card) == 0 {
		return nil, errors.New("Информация о банковской карте отсутствует")
	}

	return user.Card, nil
}

func (u *User) Delete() error {
	db := database.GetDB()

	return db.Select("Card", "Company", "Donations.Wards").Delete(&u).Error
}

func (u *User) DeleteByID() error {
	dbase := database.GetDB()
	result := dbase.Select("Card", "Company", "Donations.Wards").Delete(&User{ID: u.ID})

	return result.Error
}

func (u *User) Update() error {
	dbase := database.GetDB()
	result := dbase.Updates(&u)
	return result.Error
}
