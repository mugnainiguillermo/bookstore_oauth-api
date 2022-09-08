package rest

import (
	"github.com/go-resty/resty/v2"
	"github.com/mugnainiguillermo/bookstore_oauth-api/src/domain/user"
	"github.com/mugnainiguillermo/bookstore_utils-go/rest_errors"
	"time"
)

var (
	client *resty.Client
)

func init() {
	client = resty.New().
		SetBaseURL("http://localhost:9000").
		SetTimeout(30 * time.Second)
}

type RestUsersRepository interface {
	LoginUser(string, string) (*user.User, *rest_errors.RestErr)
}

type usersRepository struct {
}

func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

type Result struct {
	Msg string `json:"msg"`
}

func (r *usersRepository) LoginUser(email string, password string) (*user.User, *rest_errors.RestErr) {
	request := user.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	var user user.User
	var restErr rest_errors.RestErr

	resp, err := client.R().
		SetBody(request).
		SetResult(&user).
		SetError(&restErr).
		Post("/users/login")

	if err != nil {
		//TODO: Log properly
		return nil, rest_errors.NewInternalServerError("error during client request", err)
	}

	if resp.IsError() {
		return nil, &restErr
	}

	return &user, nil
}
