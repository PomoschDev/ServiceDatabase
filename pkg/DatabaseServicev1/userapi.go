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
	"strings"
	"time"
)

func (s *serverAPI) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	newUser := new(models.User)

	err := utilities.Transformation(req, newUser)
	if err != nil {
		logger.Log.Error("", sl.Err(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Ошибка при конвертации модели: %v", err))
	}

	logger.Log.Info("конвертация: ", slog.Any("model", newUser), slog.Any("req", req))

	newUser.CreatedAt = time.Now()
	newUser.UpdatedAt = time.Now()

	if err := newUser.Create(); err != nil {
		logger.Log.Error("newUser.Create()", sl.Err(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Ошибка при создании пользователя: %v", err))
	}

	db := database.GetDB()
	findUser := new(models.User)
	result := db.Preload("Card").Preload("Company").Preload("Donations").Where(&models.User{Phone: newUser.Phone}).
		Find(&findUser)

	if result.Error != nil {
		logger.Log.Error("findUser", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", result.Error))
	}

	if findUser.ID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Созданный пользователь не найден: %v", result.Error))
	}

	response := &CreateUserResponse{}

	err = utilities.Transformation(findUser, response)
	if err != nil {
		logger.Log.Error("utilities.Transformation(findUser, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	logger.Log.Info("response: ", slog.Any("", response))

	return response, nil
}

func (s *serverAPI) Users(ctx context.Context, req *Empty) (*UsersResponse, error) {
	users, err := models.AllUsers()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	obj := struct {
		Users []*models.User `json:"users,omitempty"`
	}{}

	obj.Users = users

	response := &UsersResponse{}

	err = utilities.Transformation(obj, response)
	if err != nil {
		logger.Log.Error("utilities.Transformation(obj, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) IsRole(ctx context.Context, req *IsRoleRequest) (*IsRoleResponse, error) {
	user := &models.User{ID: req.GetId()}

	accessory, err := user.IsRole(req.GetRole())
	if err != nil {
		logger.Log.Error("accessory, err := user.IsRole(req.GetRole())", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	logger.Log.Info("Модель: ", slog.Any("user", user))

	response := &IsRoleResponse{Accessory: accessory}

	return response, nil
}

func (s *serverAPI) ComparePassword(ctx context.Context, req *ComparePasswordRequest) (*ComparePasswordResponse, error) {
	user := &models.User{Phone: req.GetPhone()}
	if err := user.FindUserPhone(); err != nil {
		logger.Log.Error("user.FindUserPhone()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if user.ID == 0 {
		logger.Log.Warn("user.ID == 0")
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Пользователь не найден"))
	}

	isValid, err := models.ComparePassword(req.GetPhone(), req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка при сравнении пароля: %v", err))
	}

	response := &ComparePasswordResponse{
		Accessory: isValid,
	}

	return response, nil
}

func (s *serverAPI) UserIsExists(ctx context.Context, req *UserIsExistsRequest) (*UserIsExistsResponse, error) {
	if len(req.GetPhone()) == 0 || req.GetPhone() == "" {
		logger.Log.Warn("len(req.GetPhone()) == 0 req.GetPhone() == \"\"")
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Неверные аргументы"))
	}

	isExists, err := models.UserIsExists(req.GetPhone())
	if err != nil {
		logger.Log.Error("models.UserIsExists(req.GetPhone())", sl.Err(err))
		return nil, status.Error(codes.NotFound, fmt.Sprintf("%v", err))
	}

	return &UserIsExistsResponse{IsExists: isExists}, nil
}

func (s *serverAPI) FindUserById(ctx context.Context, req *FindUserByIdRequest) (*CreateUserResponse, error) {
	if req.GetId() <= 0 {
		logger.Log.Warn("req.GetId() == 0")
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("ID не может быть меньше или равен 0"))
	}

	user := &models.User{ID: req.GetId()}
	if err := user.FindUserID(); err != nil {
		logger.Log.Error("user.FindUserID()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if user.Phone == "" || len(user.Phone) == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Пользователь с id %d не найден", req.GetId()))
	}

	response := new(CreateUserResponse)

	if err := utilities.Transformation(user, response); err != nil {
		logger.Log.Error("utilities.Transformation(user, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) FindUserByEmail(ctx context.Context, req *FindUserByEmailRequest) (*CreateUserResponse, error) {
	if len(req.GetEmail()) == 0 || req.GetEmail() == "" {
		logger.Log.Warn("len(req.GetEmail()) == 0 || req.GetEmail() == \"\"")
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Email не может быть пустым"))
	}
	user := &models.User{Email: req.GetEmail()}
	if err := user.FindUserEmail(); err != nil {
		logger.Log.Error("user.FindUserEmail()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if user.Email == "" || len(user.Email) == 0 || user.ID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Пользователь с email %s не найден", req.GetEmail()))
	}

	response := new(CreateUserResponse)

	if err := utilities.Transformation(user, response); err != nil {
		logger.Log.Error("utilities.Transformation(user, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) FindUserByPhone(ctx context.Context, req *FindUserByPhoneRequest) (*CreateUserResponse, error) {
	if len(req.GetPhone()) == 0 || req.GetPhone() == "" {
		logger.Log.Warn("len(req.GetPhone()) == 0 || req.GetPhone() == \"\"")
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Неверные аргументы"))
	}
	user := &models.User{Phone: req.GetPhone()}
	if err := user.FindUserPhone(); err != nil {
		logger.Log.Error("user.FindUserPhone()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if user.ID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Пользователь с номером телефона %s не найден",
			req.GetPhone()))
	}

	response := new(CreateUserResponse)

	if err := utilities.Transformation(user, response); err != nil {
		logger.Log.Error("utilities.Transformation(user, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) ChangeUserType(ctx context.Context, req *ChangeUserTypeRequest) (*ChangeUserTypeResponse, error) {
	if req.GetId() <= 0 {
		logger.Log.Error("req.GetId() <= 0")
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("ID не может быть меньше или равен 0"))
	}

	if req.GetType() > 1 {
		logger.Log.Error("req.GetType() > 1")
		return nil, status.Error(codes.InvalidArgument,
			fmt.Sprintf("Поле type может принимать значения только 0 или 1"))
	}

	user := &models.User{ID: req.GetId()}
	if err := user.ChangeUserType(req.GetType()); err != nil {
		logger.Log.Error("user.ChangeUserType(req.GetType())", sl.Err(err))
		return nil, status.Error(codes.NotFound, fmt.Sprintf("%v", err))
	}

	response := &ChangeUserTypeResponse{
		Accessory: true,
	}

	return response, nil
}

func (s *serverAPI) FindUserCompany(ctx context.Context, req *FindUserCompanyRequest) (*Company, error) {
	if req.GetId() <= 0 {
		logger.Log.Error("req.GetId() <= 0")
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("ID не может быть меньше или равен 0"))
	}

	company, err := models.FindUserCompany(req.GetId())
	if err != nil {
		logger.Log.Error("models.FindUserCompany(req.GetId())", sl.Err(err))
		return nil, status.Error(codes.NotFound, fmt.Sprintf("%v", err))
	}

	response := new(Company)

	if err := utilities.Transformation(company, response); err != nil {
		logger.Log.Error("utilities.Transformation(company, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) FindUserDonations(ctx context.Context,
	req *FindUserDonationsRequest) (*FindUserDonationsResponse, error) {
	if req.GetId() <= 0 {
		logger.Log.Error("req.GetId() <= 0")
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("ID не может быть меньше или равен 0"))
	}

	donations, err := models.FindUserDonations(req.GetId())
	if err != nil {
		logger.Log.Error("models.FindUserDonations(req.GetId())", sl.Err(err))
		return nil, status.Error(codes.NotFound, fmt.Sprintf("%v", err))
	}

	obj := struct {
		Donations []*models.Donations `json:"donations"`
	}{}

	obj.Donations = donations

	response := new(FindUserDonationsResponse)

	if err := utilities.Transformation(obj, response); err != nil {
		logger.Log.Error("utilities.Transformation(company, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) FindUserCard(ctx context.Context, req *FindUserCardRequest) (*FindUserCardResponse, error) {
	if req.GetId() <= 0 {
		logger.Log.Error("req.GetId() <= 0")
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("ID не может быть меньше или равен 0"))
	}

	cards, err := models.FindUserCard(req.GetId())
	if err != nil && !strings.Contains(err.Error(), "Информация о банковской карте отсутствует") {
		logger.Log.Error("models.FindUserCard(req.GetId())", sl.Err(err))
		return nil, status.Error(codes.NotFound, fmt.Sprintf("%v", err))
	}

	if err != nil && strings.Contains(err.Error(), "Информация о банковской карте отсутствует") {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("%v", err))
	}

	obj := struct {
		Cards []*models.Card `json:"cards"`
	}{}

	obj.Cards = cards

	response := new(FindUserCardResponse)

	if err := utilities.Transformation(obj, response); err != nil {
		logger.Log.Error("utilities.Transformation(obj, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) DeleteUserByModel(ctx context.Context, req *DeleteUserByModelRequest) (*HTTPCodes, error) {
	user := new(models.User)
	if err := utilities.Transformation(req.User, user); err != nil {
		logger.Log.Error("utilities.Transformation(req, obj)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if err := user.Delete(); err != nil {
		logger.Log.Error("user.Delete()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := &HTTPCodes{Code: http.StatusOK}

	return response, nil
}

func (s *serverAPI) DeleteUserById(ctx context.Context, req *DeleteUserByIdRequest) (*HTTPCodes, error) {
	if req.GetId() <= 0 {
		logger.Log.Error("req.GetId() <= 0")
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("ID не может быть меньше или равен 0"))
	}

	user := &models.User{ID: req.GetId()}
	if err := user.DeleteByID(); err != nil {
		logger.Log.Error("user.DeleteByID()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := &HTTPCodes{Code: http.StatusOK}

	return response, nil
}

func (s *serverAPI) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*CreateUserResponse, error) {
	user := new(models.User)

	if err := utilities.Transformation(req, user); err != nil {
		logger.Log.Error("utilities.Transformation(req, user)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if err := user.Update(); err != nil {
		logger.Log.Error("user.Update()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	db := database.GetDB()
	user = nil
	result := db.Preload("Card").Preload("Company").Preload("Donations").Where(&models.User{ID: req.Id}).
		Find(&user)

	if result.Error != nil {
		logger.Log.Error("findUser", sl.Err(result.Error))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", result.Error))
	}

	if user.ID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Обновленный пользователь не найден: %v", result.Error))
	}

	response := new(CreateUserResponse)
	if err := utilities.Transformation(user, response); err != nil {
		logger.Log.Error("utilities.Transformation(req, user)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) AddCardToUser(ctx context.Context, req *AddCardToUserRequest) (*AddCardToUserResponse, error) {
	user := new(models.User)
	card := new(models.Card)

	if err := utilities.Transformation(req.GetCard(), card); err != nil {
		logger.Log.Error("utilities.Transformation(req.GetCard(), card)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	db := database.GetDB()
	user = nil
	result := db.Preload("Card").Preload("Company.Card").Preload("Donations.Ward").Where(&models.User{ID: card.
		UserID}).
		Find(&user)

	if result.Error != nil {
		logger.Log.Error("findUser", sl.Err(result.Error))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", result.Error))
	}

	if user.ID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Обновленный пользователь не найден: %v", result.Error))
	}

	user.Card = append(user.Card, card)

	if err := user.Update(); err != nil {
		logger.Log.Error("user.Update()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := &AddCardToUserResponse{}
	if err := utilities.Transformation(user, response); err != nil {
		logger.Log.Error("utilities.Transformation(req, user)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}
