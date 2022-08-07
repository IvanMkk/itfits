package application

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (a *App) RegisterUser(username string, password string) (string, error) {

	queryString := `insert into it_users(
		user_id,
		email,
		status,
		password,
		dtime_created)
	values ($1, $2, $3, $4, $5)`
	stmt, err := a.DB.Prepare(queryString)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	userId := uuid.New()
	dt := time.Now()
	timeCreated := dt.Format(timeLayout)

	_, err = stmt.Exec(
		userId,
		username,
		"active",
		hashedPassword,
		timeCreated)

	if err != nil {
		return "", err
	}

	return "Success", nil
}
