package application

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (a *App) GenerateToken(username string, password string) (map[string]interface{}, error) {

	queryString := "select user_id, password from users where email = $1"

	stmt, err := a.DB.Prepare(queryString)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	userId := ""
	accountPassword := ""
	err = stmt.QueryRow(username).Scan(&userId, &accountPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Invalid username or password.")
		}

		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(accountPassword), []byte(password))
	if err != nil {
		return nil, errors.New("Invalid username or password.")
	}
	queryString = "insert into authentication_tokens(token_id, user_id, auth_token, generated_at, expires_at) values ($1, $2, $3, $4, $5)"
	stmt, err = a.DB.Prepare(queryString)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	randomToken := make([]byte, 32)
	_, err = rand.Read(randomToken)
	if err != nil {
		return nil, err
	}

	authToken := base64.URLEncoding.EncodeToString(randomToken)
	const timeLayout = "2006-01-02 15:04:05"
	dt := time.Now()
	expirtyTime := time.Now().Add(time.Minute * 1)

	generatedAt := dt.Format(timeLayout)
	expiresAt := expirtyTime.Format(timeLayout)

	token_id := uuid.New()

	_, err = stmt.Exec(token_id, userId, authToken, generatedAt, expiresAt)

	if err != nil {
		return nil, err
	}

	tokenDetails := map[string]interface{}{
		"token_type":   "Bearer",
		"auth_token":   authToken,
		"generated_at": generatedAt,
		"expires_at":   expiresAt,
	}

	return tokenDetails, nil
}

func (a *App) ValidateToken(authToken string) (map[string]interface{}, error) {

	queryString := `select 
                users.user_id,
                email,
                generated_at,
                expires_at                         
            from authentication_tokens
            left join users
            on authentication_tokens.user_id = users.user_id
            where auth_token = ?`

	stmt, err := a.DB.Prepare(queryString)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	userId := ""
	username := ""
	generatedAt := ""
	expiresAt := ""

	err = stmt.QueryRow(authToken).Scan(&userId, &username, &generatedAt, &expiresAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Invalid access token.")
		}

		return nil, err
	}

	const timeLayout = "2006-01-02 15:04:05"
	expiryTime, _ := time.Parse(timeLayout, expiresAt)
	currentTime, _ := time.Parse(timeLayout, time.Now().Format(timeLayout))

	if expiryTime.Before(currentTime) {
		return nil, errors.New("The token is expired.")
	}

	// TODO: change it to JSON as everywere
	userDetails := map[string]interface{}{
		"user_id":      userId,
		"username":     username,
		"generated_at": generatedAt,
		"expires_at":   expiresAt,
	}

	return userDetails, nil
}
