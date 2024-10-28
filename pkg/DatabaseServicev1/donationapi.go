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

func (s *serverAPI) Donations(ctx context.Context, req *Empty) (*DonationsResponse, error) {
	donations, err := models.AllDonations()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := new(DonationsResponse)

	err = utilities.Transformation(donations, &response.Donations)
	if err != nil {
		logger.Log.Error("utilities.Transformation(donations, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) CreateDonations(ctx context.Context, req *CreateDonationsRequest) (*CreateDonationsResponse, error) {
	newDonation := new(models.Donations)

	err := utilities.Transformation(req, newDonation)
	if err != nil {
		logger.Log.Error("", sl.Err(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Ошибка при конвертации модели: %v", err))
	}

	newDonation.CreatedAt = time.Now()
	newDonation.UpdatedAt = time.Now()

	if err := newDonation.Create(); err != nil {
		logger.Log.Error("newDonation.Create()", sl.Err(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Ошибка при создании доната: %v", err))
	}

	db := database.GetDB()
	findDonation := new(models.Donations)
	result := db.Preload("Ward").Where(&models.Donations{ID: newDonation.ID}).
		Find(&findDonation)

	if result.Error != nil {
		logger.Log.Error("findDonation", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", result.Error))
	}

	if findDonation.ID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Созданный донат не найден: %v", result.Error))
	}

	response := new(CreateDonationsResponse)

	err = utilities.Transformation(findDonation, response)
	if err != nil {
		logger.Log.Error("utilities.Transformation(findDonation, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	logger.Log.Info("response: ", slog.Any("", response))

	return response, nil
}

func (s *serverAPI) FindDonationWards(ctx context.Context, req *FindDonationWardsRequest) (*FindDonationWardsResponse, error) {
	wards, err := models.FindDonationWards(req.GetId())
	if err != nil {
		logger.Log.Error("models.FindDonationWards(req.GetId())", sl.Err(err))
		return nil, status.Error(codes.NotFound, fmt.Sprintf("%v", err))
	}

	response := new(FindDonationWardsResponse)

	if err := utilities.Transformation(wards, &response.Wards); err != nil {
		logger.Log.Error("utilities.Transformation(wards, response.Wards)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) FindDonationUser(ctx context.Context, req *FindDonationUserRequest) (*FindDonationUserResponse,
	error) {
	user, err := models.FindDonationUser(req.GetId())
	if err != nil {
		logger.Log.Error("models.FindDonationUser", sl.Err(err))
		return nil, status.Error(codes.NotFound, fmt.Sprintf("%v", err))
	}

	response := new(FindDonationUserResponse)

	if err := utilities.Transformation(user, &response.User); err != nil {
		logger.Log.Error("utilities.Transformation(user, &response.User)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) FindDonationById(ctx context.Context, req *FindDonationByIdRequest) (*CreateDonationsResponse, error) {
	donation := &models.Donations{ID: req.GetId()}
	if err := donation.FindDonationID(); err != nil {
		logger.Log.Error("donation.FindDonationID()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if donation.Title == "" || len(donation.Title) == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Донат с id %d не найден", req.GetId()))
	}

	response := new(CreateDonationsResponse)

	if err := utilities.Transformation(donation, response); err != nil {
		logger.Log.Error("utilities.Transformation(donation, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) DeleteDonationByModel(ctx context.Context, req *DeleteDonationByModelRequest) (*HTTPCodes, error) {
	donation := new(models.Donations)
	if err := utilities.Transformation(req, donation); err != nil {
		logger.Log.Error("utilities.Transformation(req, donation)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if err := donation.Delete(); err != nil {
		logger.Log.Error("donation.Delete()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := &HTTPCodes{Code: http.StatusOK}

	return response, nil
}

func (s *serverAPI) DeleteDonationById(ctx context.Context, req *DeleteDonationByIdRequest) (*HTTPCodes, error) {
	donation := &models.Donations{ID: req.GetId()}
	if err := donation.DeleteByID(); err != nil {
		logger.Log.Error("donation.DeleteByID()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := &HTTPCodes{Code: http.StatusOK}

	return response, nil
}

func (s *serverAPI) UpdateDonation(ctx context.Context, req *UpdateDonationsRequest) (*CreateDonationsResponse, error) {
	donation := new(models.Donations)

	if err := utilities.Transformation(req, donation); err != nil {
		logger.Log.Error("utilities.Transformation(req, donation)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if err := donation.Update(); err != nil {
		logger.Log.Error("donation.Update()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	db := database.GetDB()
	donation = nil
	result := db.Preload("Wards").Where(&models.Donations{ID: req.Id}).
		Find(&donation)

	if result.Error != nil {
		logger.Log.Error("findDonation", sl.Err(result.Error))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", result.Error))
	}

	if donation.ID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Обновленный донат не найден: %v", result.Error))
	}

	response := &CreateDonationsResponse{}
	if err := utilities.Transformation(donation, response); err != nil {
		logger.Log.Error("utilities.Transformation(donation, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}
