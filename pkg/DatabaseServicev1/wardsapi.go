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
	"net/http"
	"time"
)

func (s *serverAPI) Wards(ctx context.Context, req *Empty) (*WardsResponse, error) {
	wards, err := models.AllWards()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := new(WardsResponse)

	err = utilities.Transformation(wards, &response.Wards)
	if err != nil {
		logger.Log.Error("utilities.Transformation(wards, response.Wards)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) CreateWard(ctx context.Context, req *CreateWardRequest) (*Ward, error) {
	newWard := new(models.Ward)

	err := utilities.Transformation(req, newWard)
	if err != nil {
		logger.Log.Error("", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка при конвертации модели: %v", err))
	}

	newWard.CreatedAt = time.Now()
	newWard.UpdatedAt = time.Now()

	if err := newWard.Create(); err != nil {
		logger.Log.Error("newWard.Create()", sl.Err(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Ошибка при создании пользователя: %v", err))
	}

	db := database.GetDB()
	findWard := new(models.Ward)
	result := db.Where(&models.Ward{ID: newWard.ID}).
		Find(&findWard)

	if result.Error != nil {
		logger.Log.Error("findWard", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", result.Error))
	}

	if findWard.ID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Созданный подопечный не найден: %v", result.Error))
	}

	response := new(Ward)

	err = utilities.Transformation(findWard, response)
	if err != nil {
		logger.Log.Error("utilities.Transformation(findWard, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) FindWardById(ctx context.Context, req *FindWardByIdRequest) (*Ward, error) {
	ward := &models.Ward{ID: req.GetId()}
	if err := ward.FindWardID(); err != nil {
		logger.Log.Error("ward.FindWardID()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if ward.Title == "" || len(ward.Title) == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Подопечный с id %d не найден", req.GetId()))
	}

	response := new(Ward)

	if err := utilities.Transformation(ward, response); err != nil {
		logger.Log.Error("utilities.Transformation(ward, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) DeleteWardByModel(ctx context.Context, req *Ward) (*HTTPCodes, error) {
	ward := new(models.Ward)
	if err := utilities.Transformation(req, ward); err != nil {
		logger.Log.Error("utilities.Transformation(req, ward)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if err := ward.Delete(); err != nil {
		logger.Log.Error("ward.Delete()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := &HTTPCodes{Code: http.StatusOK}

	return response, nil
}

func (s *serverAPI) DeleteWardById(ctx context.Context, req *DeleteWardByIdRequest) (*HTTPCodes, error) {
	ward := &models.Ward{ID: req.GetId()}
	if err := ward.DeleteByID(); err != nil {
		logger.Log.Error("ward.DeleteByID()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := &HTTPCodes{Code: http.StatusOK}

	return response, nil
}

func (s *serverAPI) UpdateWard(ctx context.Context, req *Ward) (*Ward, error) {
	ward := new(models.Ward)

	if err := utilities.Transformation(req, ward); err != nil {
		logger.Log.Error("utilities.Transformation(req, ward)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if err := ward.Update(); err != nil {
		logger.Log.Error("ward.Update()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	db := database.GetDB()
	ward = nil
	result := db.Where(&models.Ward{ID: req.Id}).
		Find(&ward)

	if result.Error != nil {
		logger.Log.Error("findWard", sl.Err(result.Error))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", result.Error))
	}

	if ward.ID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Обновленный подопечный не найден: %v", result.Error))
	}

	response := new(Ward)
	if err := utilities.Transformation(ward, response); err != nil {
		logger.Log.Error("utilities.Transformation(ward, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}
