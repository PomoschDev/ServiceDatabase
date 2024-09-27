package DatabaseServicev1

import (
	"DababaseService/pkg/config"
	"DababaseService/pkg/database"
	"DababaseService/pkg/database/models"
	"DababaseService/pkg/utilities"
	"testing"
	"time"
)

func preload() error {
	cfg := config.MustLoadByPath("../../config/local.yaml")
	_, err := database.Init(cfg)

	return err
}

func Test_Transformation(t *testing.T) {
	cr := &CreateUserRequest{}
	err := preload()
	if err != nil {
		t.Errorf("Ошибка preload: %v", err)
	}

	ur := &models.User{
		ID:       1,
		Email:    "1",
		Username: "1",
		Password: "1",
		Phone:    "1",
		Card: []*models.Card{&models.Card{
			ID:        2,
			FullName:  "asd",
			Number:    "asd",
			Date:      "ads",
			Cvv:       123,
			UserID:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}},
		Role:      "1",
		Company:   nil,
		Type:      0,
		Donations: nil,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = utilities.Transformation(ur, cr)
	if err != nil {
		t.Errorf("Ошибка при трансформации: %v", err)
		return
	}

	newur := &models.User{}

	err = utilities.Transformation(cr, newur)
	if err != nil {
		t.Errorf("Ошибка при трансформации: %v", err)
		return
	}

	t.Logf("Результат1: %+v", cr)
	err = utilities.Transformation(cr, newur)
	if err != nil {
		t.Errorf("Ошибка при трансформации: %v", err)
		return
	}
	t.Logf("Результат2: %+v", newur)
}
