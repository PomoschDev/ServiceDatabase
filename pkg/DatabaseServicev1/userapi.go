package DatabaseServicev1

import (
	"DababaseService/iternal/lib/logger/sl"
	"DababaseService/pkg/database"
	"DababaseService/pkg/database/models"
	"DababaseService/pkg/logger"
	"DababaseService/pkg/paths"
	"DababaseService/pkg/utilities"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
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

	response := new(UsersResponse)

	err = utilities.Transformation(users, &response.Users)
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
		if strings.Contains(err.Error(), "Пользователь не найден") {
			logger.Log.Error("accessory, err := user.IsRole(req.GetRole())", sl.Err(err))
			return nil, status.Error(codes.NotFound, fmt.Sprintf("%v", err))
		}
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

	response := new(FindUserDonationsResponse)

	if err := utilities.Transformation(donations, &response.Donations); err != nil {
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

	response := new(FindUserCardResponse)

	if err := utilities.Transformation(cards, &response.Cards); err != nil {
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

func (s *serverAPI) SetUserAvatar(stream DatabaseService_SetUserAvatarServer) error {
	chErr := make(chan error, 1)
	data := make([]byte, 0, 5_000_000)
	userId := uint64(0)
	imageType := ""
	defer stream.Context().Done()

	go func() {
		defer close(chErr)
	loop:
		for {
			req := new(SetUserAvatarRequest)
			err := stream.RecvMsg(req)
			if err == io.EOF {
				chErr <- nil
				break loop
			}
			if err != nil {
				logger.Log.Error("SetUserAvatar", sl.Err(err))
				chErr <- status.Error(codes.DataLoss, "Не удалось прочитать изображение")
			}

			switch u := req.Data.(type) {
			case *SetUserAvatarRequest_UserId:
				userId = u.UserId
			case *SetUserAvatarRequest_ImageType:
				imageType = u.ImageType
			case *SetUserAvatarRequest_ChunkData:
				data = append(data, u.ChunkData...)
			}
		}
	}()

	err := <-chErr
	if err != nil {
		logger.Log.Error("Ошибка после цикла", sl.Err(err))
		return err
	}

	//Сохраняем в файл
	{
		//Ищем старые файлы
		{
			pathFiles, err := findFilePath(utilities.MD5(fmt.Sprintf("user_%d", userId)))
			if err != nil {
				logger.Log.Error("Ошибка при поиске файлов: ", sl.Err(err))
				return status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
			}

			logger.Log.Info("Нашли файлы", slog.Any("paths", pathFiles))

			for _, i := range pathFiles {
				os.Remove(filepath.Join(paths.IMAGE_DIR, i))
			}
		}

		fileName := fmt.Sprintf("%s.%s", utilities.MD5(fmt.Sprintf("user_%d", userId)), imageType)
		pathToFile := filepath.Join(paths.IMAGE_DIR, fileName)
		f, err := os.OpenFile(pathToFile, os.O_CREATE|os.O_APPEND|os.O_TRUNC, os.ModePerm)
		if err != nil {
			logger.Log.Error("Ошибка при открытии файла", sl.Err(err))
			return status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
		}
		defer f.Close()

		_, err = f.Write(data)
		if err != nil {
			logger.Log.Error("Ошибка при записи файла: %v", err)
			return status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
		}

	}

	response := new(HTTPCodes)
	response.Code = http.StatusOK

	err = stream.SendMsg(response)
	if err != nil {
		logger.Log.Error("Ошибка при закрытии стрима", sl.Err(err))
		return status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	return nil
}

func findFilePath(fileName string) ([]string, error) {
	images, err := fs.Glob(os.DirFS(paths.IMAGE_DIR), fmt.Sprintf("%s*", fileName))
	if err != nil {
		logger.Log.Error("DeleteUserAvatar", sl.Err(err))
		return nil, err
	}

	return images, nil
}

func (s *serverAPI) DeleteUserAvatar(ctx context.Context, req *DeleteUserAvatarRequest) (*HTTPCodes, error) {
	fileName := utilities.MD5(fmt.Sprintf("user_%d", req.UserId))
	images, err := fs.Glob(os.DirFS(paths.IMAGE_DIR), fmt.Sprintf("%s*", fileName))
	if err != nil {
		logger.Log.Error("DeleteUserAvatar", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	for _, i := range images {
		logger.Log.Info("i", slog.String("name", i))
		fileName = i
	}

	pathToFile := filepath.Join(paths.IMAGE_DIR, fileName)

	err = os.Remove(pathToFile)
	if err != nil {
		logger.Log.Error("os.Remove(pathToFile)", sl.Err(err))
		return nil, status.Error(codes.Internal, fmt.Sprintf("Ошибка при удалении файла: %v", err))
	}

	request := &HTTPCodes{Code: http.StatusOK}

	return request, nil
}

func (s *serverAPI) GetUserAvatar(req *GetUserAvatarRequest, stream DatabaseService_GetUserAvatarServer) error {
	user := &models.User{ID: req.GetUserId()}
	const chunkSize = 1024 // Размер части в байтах
	defer stream.Context().Done()

	if err := user.FindUserID(); err != nil {
		logger.Log.Error("FindUserID", sl.Err(err))
		return status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервера: %v", err))
	}

	if user.Password == "" || len(user.Password) == 0 {
		logger.Log.Error("Пользователь не найден", slog.String("password", user.Password))
		return status.Error(codes.NotFound, "Пользователь не найден")
	}

	fileName := utilities.MD5(fmt.Sprintf("user_%d", user.ID))
	images, err := fs.Glob(os.DirFS(paths.IMAGE_DIR), fmt.Sprintf("%s*", fileName))
	if err != nil {
		logger.Log.Error("GetUserAvatar", sl.Err(err))
		return status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	for _, i := range images {
		fileName = i
	}

	pathToFile := filepath.Join(paths.IMAGE_DIR, fileName)

	f, err := os.Open(pathToFile)
	if err != nil {
		logger.Log.Error("os.Open", sl.Err(err))
		return status.Error(codes.NotFound, fmt.Sprintf("Изображение отсутствует"))
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		logger.Log.Error("Ошибка при чтении Stat:", sl.Err(err))
		return status.Error(codes.Internal, fmt.Sprintf("Ошибка на стороне сервиса: %v", err))
	}

	data := make([]byte, fileInfo.Size()/chunkSize)

	imageInfo := &GetUserAvatarResponse{Data: &GetUserAvatarResponse_Info{
		Info: &ImageInfo{
			FileName: fileName,
			Type:     strings.Replace(filepath.Ext(fileName), ".", "", -1),
			Size:     fileInfo.Size(),
		},
	}}

	err = stream.SendMsg(imageInfo)
	if err != nil {
		logger.Log.Error("Ошибка при отправке imageInfo", sl.Err(err))
		return status.Error(codes.Internal, fmt.Sprintf("Ошибка при отправке информации об изображении: %v", err))
	}

	done := make(chan struct{}, 1)
	chErr := make(chan error, 1)

	go func() {
		defer close(done)
		defer close(chErr)
	loop:
		for {
			n, err := f.Read(data)
			if err == io.EOF {
				chErr <- nil
				done <- struct{}{}
				break loop
			}

			if err != nil {
				done <- struct{}{}
				chErr <- status.Error(codes.Internal, fmt.Sprintf("Ошибка при чтении файла: %v", err))
			}

			data = data[:n]
			reqChunk := &GetUserAvatarResponse{
				Data: &GetUserAvatarResponse_ChunkData{ChunkData: data},
			}

			err = stream.SendMsg(reqChunk)
			if err != nil && err != io.EOF {
				done <- struct{}{}
				chErr <- status.Error(codes.Internal, fmt.Sprintf("Ошибка при отправке сообщения в поток: %v", err))
			}
		}
	}()

	err = <-chErr
	<-done

	if err != nil {
		return err
	}

	return nil
}
