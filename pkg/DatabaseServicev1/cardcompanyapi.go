package DatabaseServicev1

import (
	"DababaseService/iternal/lib/logger/sl"
	"DababaseService/pkg/database"
	"DababaseService/pkg/database/models"
	"DababaseService/pkg/logger"
	"DababaseService/pkg/utilities"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"net/http"
	"time"
)

func (s *serverAPI) CardsCompanies(ctx context.Context, req *Empty) (*CardsCompaniesResponse, error) {
	cardCompanies, err := models.AllCardsCompanies()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := new(CardsCompaniesResponse)

	err = utilities.Transformation(cardCompanies, &response.Cards)
	if err != nil {
		logger.Log.Error("utilities.Transformation(obj, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) CreateCardCompany(ctx context.Context, req *CreateCardCompanyRequest) (*CardCompany, error) {
	newCardCompany := new(models.CardCompany)

	err := utilities.Transformation(req, newCardCompany)
	if err != nil {
		logger.Log.Error("", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка при конвертации модели: %v", err))
	}

	newCardCompany.CreatedAt = time.Now()
	newCardCompany.UpdatedAt = time.Now()

	if err := newCardCompany.Create(); err != nil {
		logger.Log.Error("newCardCompany.Create()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка при создании компании: %v", err))
	}

	db := database.GetDB()
	findCardCompany := new(models.CardCompany)
	result := db.Where(&models.CardCompany{Number: newCardCompany.Number}).
		Find(&findCardCompany)

	if result.Error != nil {
		logger.Log.Error("findCardCompany", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", result.Error))
	}

	if findCardCompany.ID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Созданная карта компании не найден"))
	}

	response := new(CardCompany)

	err = utilities.Transformation(findCardCompany, response)
	if err != nil {
		logger.Log.Error("utilities.Transformation(findCardCompany, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	logger.Log.Info("response: ", slog.Any("", response))

	return response, nil
}

func (s *serverAPI) FindCardCompanyByID(ctx context.Context, req *FindCardCompanyByIDRequest) (*CardCompany, error) {
	cardCompany := &models.CardCompany{ID: req.GetId()}
	if err := cardCompany.FindCardCompanyID(); err != nil {
		logger.Log.Error("cardCompany.FindCardCompanyID()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if cardCompany.CompanyID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Компания не найдена"))
	}

	response := new(CardCompany)

	if err := utilities.Transformation(cardCompany, response); err != nil {
		logger.Log.Error("utilities.Transformation(cardCompany, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) DeleteCardCompanyByModel(ctx context.Context, req *CardCompany) (*HTTPCodes, error) {
	cardCompany := new(models.CardCompany)
	if err := utilities.Transformation(req, cardCompany); err != nil {
		logger.Log.Error("utilities.Transformation(req, cardCompany)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if err := cardCompany.Delete(); err != nil {
		logger.Log.Error("cardCompany.Delete()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := &HTTPCodes{Code: http.StatusOK}

	return response, nil
}

func (s *serverAPI) DeleteCardCompanyById(ctx context.Context, req *DeleteCardCompanyByIdRequest) (*HTTPCodes, error) {
	cardCompany := &models.CardCompany{ID: req.GetId()}
	logger.Log.Info("DeleteCardCompanyById", slog.Any("cardCompany", cardCompany))
	if err := cardCompany.DeleteByID(); err != nil {
		logger.Log.Error("cardCompany.DeleteByID()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := &HTTPCodes{Code: http.StatusOK}

	return response, nil
}

func (s *serverAPI) UpdateCardCompany(ctx context.Context, req *CardCompany) (*CardCompany, error) {
	cardCompany := new(models.CardCompany)

	if err := utilities.Transformation(req, cardCompany); err != nil {
		logger.Log.Error("utilities.Transformation(req, cardCompany)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if err := cardCompany.Update(); err != nil {
		logger.Log.Error("cardCompany.Update()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	db := database.GetDB()
	cardCompany = nil
	result := db.Where(&models.Card{ID: req.GetId()}).
		Find(&cardCompany)

	if result.Error != nil {
		logger.Log.Error("findCardCompany", sl.Err(result.Error))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", result.Error))
	}

	if cardCompany.ID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Обновленная карта компании не найден"))
	}

	response := new(CardCompany)

	if err := utilities.Transformation(cardCompany, response); err != nil {
		logger.Log.Error("utilities.Transformation(cardCompany, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}
