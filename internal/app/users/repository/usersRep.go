package repository

import (
	"database/sql"
	"github.com/Andronovdima/AvitoChatAssignment/internal/models"
	"net/http"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(thisDB *sql.DB) *UserRepository {
	userRep := &UserRepository{
		db: thisDB,
	}
	return userRep
}

func (userRep *UserRepository) Create(user *models.User) error {
	return userRep.db.QueryRow(
		"INSERT INTO users (username) "+
			"VALUES ($1) RETURNING id",
		user.Username,
	).Scan(&user.ID)
}

func (userRep *UserRepository) IsExist(username string) bool {
	row := userRep.db.QueryRow(
		"SELECT username "+
			"FROM users "+
			"WHERE username = $1",
		username,
	)
	if row.Scan(&username) != nil {
		return false
	}

	return true

}

func (userRep *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	rows, err := userRep.db.Query(
		"SELECT id, username, created_at" +
			"FROM users ",
	)

	if err != nil {
		rerr := new(models.HttpError)
		rerr.StringErr = err.Error()
		rerr.StatusCode = http.StatusInternalServerError
		return nil, err
	}

	for rows.Next() {
		var us models.User
		err := rows.Scan(&us.ID, &us.Username, &us.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, us)
	}
	if err := rows.Close(); err != nil {
		rerr := new(models.HttpError)
		rerr.StringErr = err.Error()
		rerr.StatusCode = http.StatusInternalServerError
		return nil, err
	}

	return users, nil
}
