package service

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/Muhammadjon226/user_service/models"
)

//GetUser ...
func (us *UserService) GetUser(ctx context.Context, user *models.ID) (*models.User, error) {

	return nil, nil
}

//ListUsers ...
func (us *UserService) ListUsers(ctx context.Context, request *models.ListUserRequest) (*models.ListUserResponse, error) {

	resp, err := us.storage.User().ListUsers(request)
	if err == sql.ErrNoRows {
		log.Println(err)
		return nil, errors.New("Not found")
	}
	if err != nil {
		log.Println("err while get list of users", err)
		return nil, err
	}

	return resp, nil
}
