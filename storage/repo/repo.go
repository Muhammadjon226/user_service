package repo

import (
	pbUser "github.com/Muhammadjon226/user_service/genproto/user_service"
)

type UserStorageI interface{
	CreateUser(user *pbUser.User) (*pbUser.User, error)  
}