package tests

import (
	"DababaseService/pkg/DatabaseServicev1"
	"DababaseService/pkg/utilities"
	"DababaseService/tests/suite"
	"google.golang.org/grpc/status"
	"testing"
)

func Test_FindWardDonations_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	request := &DatabaseServicev1.FindWardDonationByIdRequest{Id: 6}

	response, err := st.ClientAPI.FindWardDonationById(ctx, request)
	if err != nil {
		statusRequest, ok := status.FromError(err)
		t.Logf("Ошибка при выполнении запроса: \ncode = %s\nmessage = %s, ok = %t\nerr = %v", statusRequest.Code().String(),
			statusRequest.Message(), ok, statusRequest.Err())
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(response))
}
