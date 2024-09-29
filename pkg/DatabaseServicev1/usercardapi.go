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

func (s *serverAPI) Cards(ctx context.Context, req *Empty) (*CardsResponse, error) {
	cards, err := models.AllCards()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := new(CardsResponse)

	err = utilities.Transformation(cards, response.Cards)
	if err != nil {
		logger.Log.Error("utilities.Transformation(cards, response.Cards)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) CreateCard(ctx context.Context, req *CreateCardRequest) (*Card, error) {
	newCard := new(models.Card)

	err := utilities.Transformation(req, newCard)
	if err != nil {
		logger.Log.Error("", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка при конвертации модели: %v", err))
	}

	newCard.CreatedAt = time.Now()
	newCard.UpdatedAt = time.Now()

	if err := newCard.Create(); err != nil {
		logger.Log.Error("newCard.Create()", sl.Err(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Ошибка при создании карты пользователя: %v", err))
	}

	db := database.GetDB()
	findCard := new(models.Card)
	result := db.Where(&models.Card{ID: newCard.ID}).
		Find(&findCard)

	if result.Error != nil {
		logger.Log.Error("findCard", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", result.Error))
	}

	if findCard.ID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Созданная карта пользователя не найден: %v",
			result.Error))
	}

	response := new(Card)

	err = utilities.Transformation(findCard, response)
	if err != nil {
		logger.Log.Error("utilities.Transformation(findCard, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	logger.Log.Info("response: ", slog.Any("", response))

	return response, nil
}

func (s *serverAPI) FindCardById(ctx context.Context, req *FindCardByIdRequest) (*Card, error) {
	card := &models.Card{ID: req.GetId()}
	if err := card.FindCardID(); err != nil {
		logger.Log.Error("card.FindCardID()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if card.Number == "" || len(card.Number) == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Карта пользователя с id %d не найден", req.GetId()))
	}

	response := new(Card)

	if err := utilities.Transformation(card, response); err != nil {
		logger.Log.Error("utilities.Transformation(card, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) DeleteCardByModel(ctx context.Context, req *Card) (*HTTPCodes, error) {
	card := new(models.Card)
	if err := utilities.Transformation(req, card); err != nil {
		logger.Log.Error("utilities.Transformation(req, card)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if err := card.Delete(); err != nil {
		logger.Log.Error("card.Delete()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := &HTTPCodes{Code: http.StatusOK}

	return response, nil
}

func (s *serverAPI) DeleteCardById(ctx context.Context, req *DeleteCardByIdRequest) (*HTTPCodes, error) {
	card := &models.Card{ID: req.GetId()}
	if err := card.DeleteByID(); err != nil {
		logger.Log.Error("card.DeleteByID()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := &HTTPCodes{Code: http.StatusOK}

	return response, nil
}

func (s *serverAPI) UpdateCard(ctx context.Context, req *Card) (*Card, error) {
	card := new(models.Card)

	if err := utilities.Transformation(req, card); err != nil {
		logger.Log.Error("utilities.Transformation(req, card)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if err := card.Update(); err != nil {
		logger.Log.Error("card.Update()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	db := database.GetDB()
	card = nil
	result := db.Where(&models.Card{ID: req.Id}).
		Find(&card)

	if result.Error != nil {
		logger.Log.Error("findCard", sl.Err(result.Error))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", result.Error))
	}

	if card.ID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Обновленная карта пользователя не найден: %v",
			result.Error))
	}

	response := new(Card)
	if err := utilities.Transformation(card, response); err != nil {
		logger.Log.Error("utilities.Transformation(card, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}
