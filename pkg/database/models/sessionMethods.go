package models

import (
	"DababaseService/pkg/database"
	"errors"
)

func AllSessions() ([]*Session, error) {
	db := database.GetDB()
	var sessions []*Session

	err := db.Find(&sessions).Error

	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (s *Session) Create() error {
	db := database.GetDB()

	return db.Create(&s).Error
}

func (s *Session) FindSessionID() error {
	db := database.GetDB()

	return db.Where(&Session{ID: s.ID}).Find(&s).Error
}

func (s *Session) FindSessionByUserID() error {
	db := database.GetDB()

	return db.Where(&Session{UserID: s.UserID}).Find(&s).Error
}

func (s *Session) ChangeRefreshTokenByID(newToken string) error {
	db := database.GetDB()

	if err := s.FindSessionID(); err != nil {
		return err
	}

	if s.ID == 0 || s.UserID == 0 {
		return errors.New("Сессия не найден")
	}

	return db.Model(&s).UpdateColumn("refresh_token", newToken).Error
}

func (s *Session) ChangeRefreshTokenByUserID(newToken string) error {
	db := database.GetDB()

	if err := s.FindSessionByUserID(); err != nil {
		return err
	}

	if s.ID == 0 {
		return errors.New("Сессия не найден")
	}

	return db.Model(&s).UpdateColumn("refresh_token", newToken).Error
}

func (s *Session) Delete() error {
	db := database.GetDB()

	return db.Delete(&s).Error
}

func (s *Session) DeleteByID() error {
	dbase := database.GetDB()

	return dbase.Delete(&Session{ID: s.ID}).Error
}

func (s *Session) DeleteByUserID() error {
	dbase := database.GetDB()

	return dbase.Delete(&Session{UserID: s.UserID}).Error
}
