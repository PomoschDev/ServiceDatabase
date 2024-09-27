package DatabaseServicev1

import (
	"DababaseService/iternal/lib/logger/sl"
	"DababaseService/pkg/database/models"
	"DababaseService/pkg/logger"
	"DababaseService/pkg/utilities"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	return nil, nil
}

func (s *serverAPI) FindCompanyById(ctx context.Context, req *FindCompanyByIdRequest) (*Company, error) {

	return nil, nil
}

func (s *serverAPI) FindCompanyByPhone(ctx context.Context, req *FindCompanyByIdPhone) (*Company, error) {

	return nil, nil
}

func (s *serverAPI) FindCompanyCard(ctx context.Context, req *FindCompanyCardRequest) (*CardCompany, error) {

	return nil, nil
}

func (s *serverAPI) DeleteCompanyByModel(ctx context.Context, req *DeleteCompanyByModelRequest) (*HTTPCodes, error) {

	return nil, nil
}

func (s *serverAPI) DeleteCompanyById(ctx context.Context, req *DeleteCompanyByIdRequest) (*HTTPCodes, error) {

	return nil, nil
}

func (s *serverAPI) UpdateCompany(ctx context.Context, req *UpdateCompanyRequest) (*HTTPCodes, error) {

	return nil, nil
}
