package rest

import (
	"encoding/json"
	"time"

	"github.com/Pawelek242/home_oauth-api/src/domain/users"
	"github.com/Pawelek242/home_oauth-api/src/utils/errors"
	"github.com/golang-restclient/rest"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://api.testpage.com",
		Timeout: 100 * time.Millisecond,
	}
	error []string
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct {
}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError(append(error, "restclient response when trying to login user"))
	}
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError(append(error, "invalid error interface when trying to login user"))
		}
		return nil, &restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {

		return nil, errors.NewInternalServerError(append(error, "Error when trying to unmarshal users response"))
	}
	return &user, nil
}
