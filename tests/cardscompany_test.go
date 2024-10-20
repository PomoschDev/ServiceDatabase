package tests

import (
	"DababaseService/pkg/utilities"
	"DababaseService/tests/suite"
	"testing"
)

func Test_CardCompany_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

	respUsers, err := st.ClientAPI.CardsCompanies(ctx, nil)
	if err != nil {
		t.Errorf("Ошибка при выполнении запроса: %v", err)
		return
	}

	t.Logf("Ответ от сервера: %s", utilities.ToJSON(respUsers))
}
