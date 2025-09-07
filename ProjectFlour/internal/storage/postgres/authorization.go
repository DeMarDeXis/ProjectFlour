package postgres

import (
	"ProjectFlour/internal/model"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type AuthorizationStorage struct {
	db *sqlx.DB
}

func NewAuthStorage(db *sqlx.DB) *AuthorizationStorage {
	return &AuthorizationStorage{
		db: db,
	}
}

func (a *AuthorizationStorage) CreateUser(user model.User) (int, error) {
	const op = "storage.postgres.authorization.CreateUser"

	var id int
	q := `INSERT INTO users (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id`
	row := a.db.QueryRow(q, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (a *AuthorizationStorage) GetUser(username string, password string) (model.User, error) {
	const op = "storage.postgres.authorization.GetUser"

	var user model.User
	q := `SELECT id, name, username, password_hash FROM users WHERE username = $1 AND password_hash = $2`

	err := a.db.QueryRow(q, username, password).Scan(&user.ID, &user.Name, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, errors.New("user not found")
		}
		return model.User{}, err
	}

	if user.ID == 0 || user.Username == "" {
		return model.User{}, errors.New("invalid user data")
	}

	return user, nil
}
