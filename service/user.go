package service

import (
	"context"
	pbUser "github.com/Muhammadjon226/user_service/genproto/user_service"
)

//GetUser ...
func (us *UserService) GetUser(ctx context.Context, user *pbUser.ID) (*pbUser.User, error){

	return nil, nil
}

//ListUsers ...
func (us *UserService) ListUsers(ctx context.Context, user *pbUser.ListUserRequest) (*pbUser.ListUserResponse, error){

	return nil, nil
}