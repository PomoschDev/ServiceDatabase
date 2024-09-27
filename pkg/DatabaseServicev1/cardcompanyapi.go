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

func (s *serverAPI) CardsCompanies(ctx context.Context, req *Empty) (*CardsCompaniesResponse, error) {
	companies, err := models.AllCompanies()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := new(CardsCompaniesResponse)

	err = utilities.Transformation(companies, response)
	if err != nil {
		logger.Log.Error("utilities.Transformation(obj, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) CreateCardCompany(ctx context.Context, req *CreateCardCompanyRequest) (*CardCompany, error) {

	return nil, nil
}

func (s *serverAPI) FindCardCompanyByID(ctx context.Context, req *FindCardCompanyByIDRequest) (*CardCompany, error) {

	return nil, nil
}

func (s *serverAPI) DeleteCardCompanyByModel(ctx context.Context, req *CardCompany) (*HTTPCodes, error) {

	return nil, nil
}

func (s *serverAPI) DeleteCardCompanyById(ctx context.Context, req *DeleteCardCompanyByIdRequest) (*HTTPCodes, error) {

	return nil, nil
}

func (s *serverAPI) UpdateCardCompany(ctx context.Context, req *CardCompany) (*CardCompany, error) {

	return nil, nil
}
