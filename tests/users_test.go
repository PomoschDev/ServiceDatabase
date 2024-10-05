package tests

import (
	"DababaseService/pkg/DatabaseServicev1"
	"DababaseService/pkg/database/models"
	"DababaseService/pkg/utilities"
	"DababaseService/tests/suite"
	"github.com/brianvoe/gofakeit/v7"
	"google.golang.org/grpc/status"
	"strconv"
	"testing"
)

const passDefaultLen = 32

func randomFakePassword() string {
	return gofakeit.Password(true, true, true, true, false, passDefaultLen)
}

func Test_CreateUser_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	user := randomUser()
	request := new(DatabaseServicev1.CreateUserRequest)

	err := utilities.Transformation(user, request)
	if err != nil {
		t.Errorf("Ошибка трансформации: %v", err)
		return
	}

	respCreateUser, err := st.ClientAPI.CreateUser(ctx, request)
	if err != nil {
		t.Errorf("Ошибка при выполнении запроса: %v", err)
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(respCreateUser))
}

func Test_Users_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	respUsers, err := st.ClientAPI.Users(ctx, nil)
	if err != nil {
		t.Errorf("Ошибка при выполнении запроса: %v", err)
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(respUsers))
}

func Test_IsRole_BadPath(t *testing.T) {
	ctx, st := suite.New(t)

	requestRandom := &DatabaseServicev1.IsRoleRequest{
		Id:   uint64(gofakeit.Uint8()),
		Role: gofakeit.JobTitle(),
	}

	_, err := st.ClientAPI.IsRole(ctx, requestRandom)
	if err != nil {
		status1, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", status1.Code().String(),
			status1.Message(), ok, status1.Err())
	}

	request := &DatabaseServicev1.IsRoleRequest{
		Id:   1,
		Role: gofakeit.JobTitle(),
	}

	respIsRole, err := st.ClientAPI.IsRole(ctx, request)
	if err != nil {
		t.Errorf("Ошибка при выполнении запроса: %v", err)
		return
	}

	t.Logf("Ответ от сервера: %t", respIsRole.Accessory)
}

func Test_ComparePassword_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	request := &DatabaseServicev1.ComparePasswordRequest{Password: randomFakePassword(), Phone: "7682740246"}
	response, err := st.ClientAPI.ComparePassword(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %t", response.Accessory)
}

func Test_ComparePassword_BadPath(t *testing.T) {
	ctx, st := suite.New(t)

	request := &DatabaseServicev1.ComparePasswordRequest{Password: randomFakePassword(), Phone: gofakeit.Phone()}
	response, err := st.ClientAPI.ComparePassword(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %t", response.Accessory)
}

func Test_UserIsExists_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	request := &DatabaseServicev1.UserIsExistsRequest{Phone: "7682740246"}
	response, err := st.ClientAPI.UserIsExists(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %t", response.IsExists)
}

func Test_UserIsExists_BadPath(t *testing.T) {
	ctx, st := suite.New(t)

	t.Logf("Проверка с пустым полем phone")
	request := &DatabaseServicev1.UserIsExistsRequest{}
	response, err := st.ClientAPI.UserIsExists(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
	}

	t.Logf("Проверка с случайным phone")
	request = &DatabaseServicev1.UserIsExistsRequest{Phone: gofakeit.Phone()}
	response, err = st.ClientAPI.UserIsExists(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %t", response.IsExists)
}

func Test_FindUserById_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	request := &DatabaseServicev1.FindUserByIdRequest{Id: 1}
	response, err := st.ClientAPI.FindUserById(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_FindUserById_BadPath(t *testing.T) {
	ctx, st := suite.New(t)

	t.Logf("Поиск с случайным id")
	request := &DatabaseServicev1.FindUserByIdRequest{Id: uint64(gofakeit.Uint8())}
	response, err := st.ClientAPI.FindUserById(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
	}

	t.Logf("Поиск с пустым id")
	request = &DatabaseServicev1.FindUserByIdRequest{}
	response, err = st.ClientAPI.FindUserById(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_FindUserByEmail_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	request := &DatabaseServicev1.FindUserByEmailRequest{Email: "tyriquepagac@borer.biz"}
	response, err := st.ClientAPI.FindUserByEmail(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_FindUserByEmail_BadPath(t *testing.T) {
	ctx, st := suite.New(t)

	t.Logf("Поиск по случайному email")
	request := &DatabaseServicev1.FindUserByEmailRequest{Email: gofakeit.Email()}
	response, err := st.ClientAPI.FindUserByEmail(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
	}

	t.Logf("Поиск по пустому email")
	request = &DatabaseServicev1.FindUserByEmailRequest{}
	response, err = st.ClientAPI.FindUserByEmail(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_FindUserByPhone_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	request := &DatabaseServicev1.FindUserByPhoneRequest{Phone: "7682740246"}
	response, err := st.ClientAPI.FindUserByPhone(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_FindUserByPhone_BadPath(t *testing.T) {
	ctx, st := suite.New(t)

	t.Logf("Поиск с пустым полем phone")
	request := &DatabaseServicev1.FindUserByPhoneRequest{}
	response, err := st.ClientAPI.FindUserByPhone(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
	}

	t.Logf("Поиск с случайным полем phone")
	request = &DatabaseServicev1.FindUserByPhoneRequest{Phone: gofakeit.Phone()}
	response, err = st.ClientAPI.FindUserByPhone(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_ChangeUserType_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	request := &DatabaseServicev1.ChangeUserTypeRequest{Id: 1, Type: 1}
	response, err := st.ClientAPI.ChangeUserType(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %t", response.Accessory)

	request = &DatabaseServicev1.ChangeUserTypeRequest{Id: 1, Type: 0}
	response, err = st.ClientAPI.ChangeUserType(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %t", response.Accessory)
}

func Test_ChangeUserType_BadPath(t *testing.T) {
	ctx, st := suite.New(t)

	t.Logf("id 0 type 1")
	request := &DatabaseServicev1.ChangeUserTypeRequest{Id: 0, Type: 1}
	response, err := st.ClientAPI.ChangeUserType(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
	}

	t.Logf("id 1 type 3")
	request = &DatabaseServicev1.ChangeUserTypeRequest{Id: 1, Type: 3}
	response, err = st.ClientAPI.ChangeUserType(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_FindUserCompany_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	request := &DatabaseServicev1.FindUserCompanyRequest{Id: 1}
	response, err := st.ClientAPI.FindUserCompany(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_FindUserCompany_BadPath(t *testing.T) {
	ctx, st := suite.New(t)

	t.Logf("Поиск с id = 0")
	request := &DatabaseServicev1.FindUserCompanyRequest{Id: 0}
	response, err := st.ClientAPI.FindUserCompany(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
	}

	t.Logf("Поиск с случайным id")
	request = &DatabaseServicev1.FindUserCompanyRequest{Id: uint64(gofakeit.Uint8())}
	response, err = st.ClientAPI.FindUserCompany(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_FindUserDonations_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	request := &DatabaseServicev1.FindUserDonationsRequest{Id: 1}
	response, err := st.ClientAPI.FindUserDonations(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_FindUserDonations_BadPath(t *testing.T) {
	ctx, st := suite.New(t)

	t.Logf("Поиск с id = 0")
	request := &DatabaseServicev1.FindUserDonationsRequest{Id: 0}
	response, err := st.ClientAPI.FindUserDonations(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
	}

	t.Logf("Поиск с случайным id")
	request = &DatabaseServicev1.FindUserDonationsRequest{Id: uint64(gofakeit.Uint8())}
	response, err = st.ClientAPI.FindUserDonations(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_FindUserCard_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	request := &DatabaseServicev1.FindUserCardRequest{Id: 6}
	response, err := st.ClientAPI.FindUserCard(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_FindUserCard_BadPath(t *testing.T) {
	ctx, st := suite.New(t)

	t.Logf("Поиск с id = 0")
	request := &DatabaseServicev1.FindUserCardRequest{Id: 0}
	response, err := st.ClientAPI.FindUserCard(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
	}

	t.Logf("Поиск с рандомным id")
	request = &DatabaseServicev1.FindUserCardRequest{Id: uint64(gofakeit.Uint8())}
	response, err = st.ClientAPI.FindUserCard(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_DeleteUserByModel_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	request := &DatabaseServicev1.FindUserByIdRequest{Id: 4}
	response, err := st.ClientAPI.FindUserById(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	requestDelete := &DatabaseServicev1.DeleteUserByModelRequest{User: response}
	responseDelete, err := st.ClientAPI.DeleteUserByModel(ctx, requestDelete)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %d", responseDelete.Code)
}

func Test_DeleteUserByModel_BadPath(t *testing.T) {
	ctx, st := suite.New(t)

	request := &DatabaseServicev1.FindUserByIdRequest{Id: 0}
	response, err := st.ClientAPI.FindUserById(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	requestDelete := &DatabaseServicev1.DeleteUserByModelRequest{User: response}
	responseDelete, err := st.ClientAPI.DeleteUserByModel(ctx, requestDelete)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %d", responseDelete.Code)
}

func Test_DeleteUserById_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	request := &DatabaseServicev1.DeleteUserByIdRequest{Id: 5}
	response, err := st.ClientAPI.DeleteUserById(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %d", response.Code)
}

func Test_DeleteUserById_BadPath(t *testing.T) {
	ctx, st := suite.New(t)

	t.Logf("Удаление с id = 0")
	request := &DatabaseServicev1.DeleteUserByIdRequest{Id: 0}
	response, err := st.ClientAPI.DeleteUserById(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
	}

	t.Logf("Удаление с случайным id")
	request = &DatabaseServicev1.DeleteUserByIdRequest{Id: uint64(gofakeit.Uint8())}
	response, err = st.ClientAPI.DeleteUserById(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
	}

	t.Logf("Ответ от сервера: %d", response.Code)
}

func Test_UpdateUser_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	requestFind := &DatabaseServicev1.FindUserByIdRequest{Id: 6}
	responseFind, err := st.ClientAPI.FindUserById(ctx, requestFind)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	responseFind.Username = "Kavon Mertz Updated"

	request := new(DatabaseServicev1.UpdateUserRequest)

	if err := utilities.Transformation(responseFind, request); err != nil {
		t.Errorf("Ошибка при конвертации: %v", err)
	}

	response, err := st.ClientAPI.UpdateUser(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_UpdateUser_BadPath(t *testing.T) {
	ctx, st := suite.New(t)

	requestFind := &DatabaseServicev1.FindUserByIdRequest{Id: uint64(gofakeit.Uint8())}
	responseFind, err := st.ClientAPI.FindUserById(ctx, requestFind)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	responseFind.Username = "Kavon Mertz Updated"

	request := new(DatabaseServicev1.UpdateUserRequest)

	if err := utilities.Transformation(responseFind, request); err != nil {
		t.Errorf("Ошибка при конвертации: %v", err)
	}

	response, err := st.ClientAPI.UpdateUser(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

func Test_AddCardToUser_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	cardData := gofakeit.CreditCard()
	cvv, err := strconv.Atoi(cardData.Cvv)
	if err != nil {
		cvv = 123
	}

	card := &DatabaseServicev1.CreateCardRequest{
		FullName: "Kavon Mertz",
		Number:   cardData.Number,
		Date:     cardData.Exp,
		Cvv:      uint64(cvv),
		UserId:   6,
	}

	request := &DatabaseServicev1.AddCardToUserRequest{Card: card}

	response, err := st.ClientAPI.AddCardToUser(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}

//--------------------

func randomUser() *models.User {
	user := new(models.User)
	user.Email = gofakeit.Email()
	user.Username = gofakeit.Name()
	user.Password = randomFakePassword()
	user.Phone = gofakeit.Phone()
	card := gofakeit.CreditCard()
	cvv, err := strconv.Atoi(card.Cvv)
	if err != nil {
		cvv = 123
	}
	user.Card = []*models.Card{
		{
			FullName: user.Username,
			Number:   card.Number,
			Date:     card.Exp,
			Cvv:      uint64(cvv),
		},
	}

	user.Role = gofakeit.JobTitle()
	user.Type = 0
	user.Donations = []*models.Donations{}

	return user
}
