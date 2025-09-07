package postgres

import (
	"ProjectFlour/internal/model"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthorizationStorage_CreateUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	dbBySQLX := sqlx.NewDb(db, "sqlmock")
	strg := NewAuthStorage(dbBySQLX)

	mock.ExpectQuery(`INSERT INTO users`).
		WithArgs("John", "John Marston", "qazwsxedc").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	id, err := strg.CreateUser(model.User{Name: "John", Username: "John Marston", Password: "qazwsxedc"})

	assert.NoError(t, err)
	assert.Equal(t, 1, id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAuthorizationStorage_GetUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	dbBySQLX := sqlx.NewDb(db, "sqlmock")
	strg := NewAuthStorage(dbBySQLX)

	mock.ExpectQuery(`SELECT id, name, username, password_hash FROM users`).
		WithArgs("not found", "pass").
		WillReturnError(sql.ErrNoRows)

	_, err := strg.GetUser("not found", "pass")
	assert.EqualError(t, err, "user not found")
}
