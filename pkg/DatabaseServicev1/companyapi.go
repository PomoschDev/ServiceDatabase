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

func (s *serverAPI) Companies(ctx context.Context, req *Empty) (*CompaniesResponse, error) {
	companies, err := models.AllCompanies()
	if err != nil {
		logger.Log.Error("models.AllCompanies()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := new(CompaniesResponse)

	if err := utilities.Transformation(companies, response); err != nil {
		logger.Log.Error("utilities.Transformation(req, user)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) CreateCompany(ctx context.Context, req *CreateCompanyRequest) (*Company, error) {
	newCompany := new(models.Company)

	err := utilities.Transformation(req, newCompany)
	if err != nil {
		logger.Log.Error("", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка при конвертации модели: %v", err))
	}

	newCompany.CreatedAt = time.Now()
	newCompany.UpdatedAt = time.Now()

	if err := newCompany.Create(); err != nil {
		logger.Log.Error("newCompany.Create()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка при создании компании: %v", err))
	}

	db := database.GetDB()
	findCompany := new(models.Company)
	result := db.Preload("Card").Where(&models.Company{Phone: newCompany.Phone}).
		Find(&findCompany)

	if result.Error != nil {
		logger.Log.Error("findCompany", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", result.Error))
	}

	if findCompany.ID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Созданная компания не найден"))
	}

	response := new(Company)

	err = utilities.Transformation(findCompany, response)
	if err != nil {
		logger.Log.Error("utilities.Transformation(findCompany, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	logger.Log.Info("response: ", slog.Any("", response))

	return response, nil
}

func (s *serverAPI) FindCompanyById(ctx context.Context, req *FindCompanyByIdRequest) (*Company, error) {
	company := &models.Company{ID: req.GetId()}
	if err := company.FindCompanyID(); err != nil {
		logger.Log.Error("company.FindCompanyID()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if company.ID == 0 || (company.Phone == "" || len(company.Phone) == 0) {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Компания не найдена"))
	}

	response := new(Company)

	if err := utilities.Transformation(company, response); err != nil {
		logger.Log.Error("utilities.Transformation(company, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) FindCompanyByPhone(ctx context.Context, req *FindCompanyByIdPhoneRequest) (*Company, error) {
	company := &models.Company{Phone: req.GetPhone()}
	if err := company.FindCompanyPhone(); err != nil {
		logger.Log.Error("company.FindCompanyPhone()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if company.ID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Компания не найдена"))
	}

	response := new(Company)

	if err := utilities.Transformation(company, response); err != nil {
		logger.Log.Error("utilities.Transformation(company, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) FindCompanyCard(ctx context.Context, req *FindCompanyCardRequest) (*CardCompany, error) {
	cardCompany, err := models.FindCompanyCard(req.GetId())
	if err != nil {
		logger.Log.Error("models.FindCompanyCard(req.GetId())", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if cardCompany.ID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("У компании отсутствует карта"))
	}

	response := new(CardCompany)

	if err := utilities.Transformation(cardCompany, response); err != nil {
		logger.Log.Error("utilities.Transformation(cardCompany, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) DeleteCompanyByModel(ctx context.Context, req *DeleteCompanyByModelRequest) (*HTTPCodes, error) {
	company := new(models.Company)
	if err := utilities.Transformation(req.Company, company); err != nil {
		logger.Log.Error("utilities.Transformation(req.Company, company)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if err := company.Delete(); err != nil {
		logger.Log.Error("company.Delete()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := &HTTPCodes{Code: http.StatusOK}

	return response, nil
}

func (s *serverAPI) DeleteCompanyById(ctx context.Context, req *DeleteCompanyByIdRequest) (*HTTPCodes, error) {
	company := &models.Company{ID: req.GetId()}
	if err := company.DeleteByID(); err != nil {
		logger.Log.Error("company.DeleteByID()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := &HTTPCodes{Code: http.StatusOK}

	return response, nil
}

func (s *serverAPI) UpdateCompany(ctx context.Context, req *UpdateCompanyRequest) (*HTTPCodes, error) {
	company := new(models.Company)

	if err := utilities.Transformation(req.GetCompany(), company); err != nil {
		logger.Log.Error("utilities.Transformation(req.Company, company)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if err := company.Update(); err != nil {
		logger.Log.Error("company.Update()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	db := database.GetDB()
	company = nil
	result := db.Preload("Card").Where(&models.Company{ID: req.GetCompany().GetId()}).
		Find(&company)

	if result.Error != nil {
		logger.Log.Error("findCompany", sl.Err(result.Error))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", result.Error))
	}

	if company.ID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Обновленная компания не найден"))
	}

	response := &HTTPCodes{Code: http.StatusOK}

	return response, nil
}
