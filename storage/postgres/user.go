package postgres

import (
	"log"

	"github.com/Muhammadjon226/user_service/models"
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

func (ur *userRepo) CreateUser(user *models.User) (*models.User, error) {

	_, err := ur.db.Exec(
		`INSERT INTO users (id, name, age) VALUES ($1, $2, $3)`, user.ID, user.Name, user.Age,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return nil, nil
}

func (ur *userRepo) ListUsers(request *models.ListUserRequest) (*models.ListUserResponse, error) {

	var (
		count int64
	)
	offset := (request.Page - 1) * request.Limit
	users := make([]*models.User, 0, request.Limit)

	err := ur.db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&count)

	if err != nil {
		return nil, err
	}

	rows, err := ur.db.Query(`SELECT * FROM users LIMIT $1 OFFSET $2`, request.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := models.User{}

		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Age,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return &models.ListUserResponse{
		Users: users,
		Count: count,
	}, nil
}
