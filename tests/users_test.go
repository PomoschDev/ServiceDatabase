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
