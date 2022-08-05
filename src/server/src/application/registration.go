package application

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (a *App) RegisterUser(username string, password string) (string, error) {

	queryString := "insert into users(user_id, email, password) values ($1, $2, $3)"
	stmt, err := a.DB.Prepare(queryString)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	userId := uuid.New()

	_, err = stmt.Exec(userId, username, hashedPassword)

	if err != nil {
		return "", err
	}

	return "Success", nil
}
