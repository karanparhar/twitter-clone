package usecase

import (
	"context"
	"testing"

	"github.com/twitter-clone/models"
	m "github.com/twitter-clone/usecase/mock"
)

func TestNewtransactionUsecase(t *testing.T) {

	r := m.NewuserRepository()
	p := m.NewtweetRepo()

	out := NewtransactionUsecase(context.TODO(), r, p)
	user := models.Profile{
		Username: "test",
		Password: "testing123",
	}
	err := out.SignUp(user)

	if err != nil {
		t.Error(err)

	}

}
