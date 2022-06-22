package repo

import "github.com/Muhammadjon226/user_service/models"

type UserStorageI interface {
	CreateUser(user *models.User) (*models.User, error)
	ListUsers(*models.ListUserRequest) (*models.ListUserResponse, error)
}
