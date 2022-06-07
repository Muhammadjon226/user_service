package postgres

import (
	"fmt"
	"log"

	pbUser "github.com/Muhammadjon226/user_service/genproto/user_service"
	"github.com/Muhammadjon226/user_service/storage/repo"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
)

type userRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{db: db}
}

func (ur *userRepo) CreateUser(user *pbUser.User) (*pbUser.User, error) {

	fmt.Println(user.Id)

	_, err := ur.db.Exec(
		`INSERT INTO users (id, name, age) VALUES ($1, $2, $3)`, user.Id, user.Name, user.Age,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return nil, nil
}
