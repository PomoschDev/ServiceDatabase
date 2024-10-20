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

func (s *serverAPI) CreateSession(ctx context.Context, req *CreateSessionRequest) (*CreateSessionResponse, error) {
	newSession := new(models.Session)

	err := utilities.Transformation(req, newSession)
	if err != nil {
		logger.Log.Error("", sl.Err(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Ошибка при конвертации модели: %v", err))
	}

	newSession.CreatedAt = time.Now()
	newSession.UpdatedAt = time.Now()

	if err := newSession.Create(); err != nil {
		logger.Log.Error("newSession.Create()", sl.Err(err))
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Ошибка при создании сессии: %v", err))
	}

	db := database.GetDB()
	findSession := new(models.Session)
	result := db.Where(&models.Session{UserID: newSession.UserID}).
		Find(&findSession)

	if result.Error != nil {
		logger.Log.Error("findSession", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", result.Error))
	}

	if findSession.ID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Созданная сессия не найдена"))
	}

	response := &CreateSessionResponse{}

	err = utilities.Transformation(findSession, response)
	if err != nil {
		logger.Log.Error("utilities.Transformation(findSession, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	logger.Log.Info("response: ", slog.Any("", response))

	return response, nil
}

func (s *serverAPI) Sessions(ctx context.Context, req *Empty) (*SessionsResponse, error) {
	sessions, err := models.AllSessions()
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := new(SessionsResponse)

	err = utilities.Transformation(sessions, &response.Sessions)
	if err != nil {
		logger.Log.Error("utilities.Transformation(obj, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) FindSessionById(ctx context.Context, req *FindSessionsByIdRequest) (*FindSessionsByIdResponse, error) {
	if req.GetId() <= 0 {
		logger.Log.Warn("req.GetId() == 0")
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("ID не может быть меньше или равен 0"))
	}

	session := &models.Session{ID: req.GetId()}
	if err := session.FindSessionID(); err != nil {
		logger.Log.Error("session.FindSessionID()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if session.UserID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Сессия с id %d не найден", req.GetId()))
	}

	response := new(FindSessionsByIdResponse)

	if err := utilities.Transformation(session, response); err != nil {
		logger.Log.Error("utilities.Transformation(session, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) FindSessionByUserId(ctx context.Context,
	req *FindSessionsByUserIdRequest) (*FindSessionsByUserIdResponse,
	error) {
	if req.GetUserId() <= 0 {
		logger.Log.Warn("req.GetUserId()")
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("UserID не может быть меньше или равен 0"))
	}

	session := &models.Session{UserID: req.GetUserId()}
	if err := session.FindSessionByUserID(); err != nil {
		logger.Log.Error("session.FindSessionByUserID()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if session.UserID == 0 {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Сессия с UserID %d не найден", req.GetUserId()))
	}

	response := new(FindSessionsByUserIdResponse)

	if err := utilities.Transformation(session, response); err != nil {
		logger.Log.Error("utilities.Transformation(session, response)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return response, nil
}

func (s *serverAPI) ChangeRefreshTokenById(ctx context.Context, req *ChangeRefreshTokenByIdRequest) (*ChangeRefreshTokenByIdResponse, error) {
	if req.GetId() <= 0 {
		logger.Log.Error("req.GetId() <= 0")
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("ID не может быть меньше или равен 0"))
	}

	if req.GetRefreshToken() == "" || len(req.GetRefreshToken()) == 0 {
		logger.Log.Error("req.GetRefreshToken() == \"\" || len(req.GetRefreshToken()) == 0")
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Refresh Token не может быть пустым"))
	}

	session := &models.Session{ID: req.GetId()}
	if err := session.ChangeRefreshTokenByID(req.GetRefreshToken()); err != nil {
		logger.Log.Error("session.ChangeRefreshTokenByID(req.GetRefreshToken())", sl.Err(err))
		return nil, status.Error(codes.NotFound, fmt.Sprintf("%v", err))
	}

	findSession := &models.Session{ID: req.GetId()}
	if err := session.FindSessionID(); err != nil {
		logger.Log.Error("session.FindSessionID()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := &ChangeRefreshTokenByIdResponse{
		Accessory: findSession.RefreshToken == req.GetRefreshToken(),
	}

	return response, nil
}

func (s *serverAPI) ChangeRefreshTokenByUserId(ctx context.Context, req *ChangeRefreshTokenByUserIdRequest) (*ChangeRefreshTokenByUserIdResponse, error) {
	if req.GetUserId() <= 0 {
		logger.Log.Error("req.GetUserId()")
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("UserID не может быть меньше или равен 0"))
	}

	if req.GetRefreshToken() == "" || len(req.GetRefreshToken()) == 0 {
		logger.Log.Error("req.GetRefreshToken() == \"\" || len(req.GetRefreshToken()) == 0")
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Refresh Token не может быть пустым"))
	}

	session := &models.Session{UserID: req.UserId}
	if err := session.ChangeRefreshTokenByUserID(req.GetRefreshToken()); err != nil {
		logger.Log.Error("session.ChangeRefreshTokenByUserID(req.GetRefreshToken())", sl.Err(err))
		return nil, status.Error(codes.NotFound, fmt.Sprintf("%v", err))
	}

	findSession := &models.Session{UserID: req.GetUserId()}
	if err := session.FindSessionByUserID(); err != nil {
		logger.Log.Error("session.FindSessionByUserID()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := &ChangeRefreshTokenByUserIdResponse{
		Accessory: findSession.RefreshToken == req.GetRefreshToken(),
	}

	return response, nil
}

func (s *serverAPI) DeleteSessionByModel(ctx context.Context, req *DeleteSessionByModelRequest) (*HTTPCodes, error) {
	session := new(models.Session)
	if err := utilities.Transformation(req, session); err != nil {
		logger.Log.Error("utilities.Transformation(req, obj)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	if err := session.Delete(); err != nil {
		logger.Log.Error("session.Delete()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := &HTTPCodes{Code: http.StatusOK}

	return response, nil
}

func (s *serverAPI) DeleteSessionById(ctx context.Context, req *DeleteSessionByIdRequest) (*HTTPCodes, error) {
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

func (s *serverAPI) DeleteSessionByUserId(ctx context.Context, req *DeleteSessionByUserIdRequest) (*HTTPCodes, error) {
	if req.GetUserId() <= 0 {
		logger.Log.Error("req.GetUserId() <= 0")
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("UserID не может быть меньше или равен 0"))
	}

	session := &models.Session{UserID: req.UserId}
	if err := session.DeleteByUserID(); err != nil {
		logger.Log.Error("session.DeleteByUserID()", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	response := &HTTPCodes{Code: http.StatusOK}

	return response, nil
}
