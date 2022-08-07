package application

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (a *App) GenerateToken(username string, password string) (map[string]interface{}, error) {

	queryString := "select user_id, password from it_users where email = $1"

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
	queryString = "insert into it_authentication_tokens(token_id, user_id, auth_token, generated_at, expires_at) values ($1, $2, $3, $4, $5)"
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

func (a *App) ValidateToken(authToken string) (map[string]string, error) {

	queryString := `select 
                it_users.user_id,
                email,
                generated_at,
                expires_at                         
            from it_authentication_tokens
            left join it_users
            on it_authentication_tokens.user_id = it_users.user_id
            where auth_token = $1`

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

	expiryTime, _ := time.Parse(timeLayoutPQ, expiresAt)
	currentTime, _ := time.Parse(timeLayout, time.Now().Format(timeLayout))

	if expiryTime.Before(currentTime) {
		return nil, errors.New("The token is expired.")
	}

	// TODO: change it to JSON as everywere
	userDetails := map[string]string{
		"user_id":      userId,
		"username":     username,
		"generated_at": generatedAt,
		"expires_at":   expiresAt,
	}

	return userDetails, nil
}

func (a *App) GetToken(c *gin.Context) (token string, err error) {
	// Get the "Authorization" header
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		return "", errors.New("Invalid Authorization header")
	}

	// Split it into two parts - "Bearer" and token
	parts := strings.SplitN(authorization, " ", 2)
	if parts[0] != "Bearer" {
		c.JSON(http.StatusBadRequest, &gin.H{
			"error": "Invalid Authorization header",
		})

		return "", errors.New("Invalid Authorization header")
	}

	// Write token into environment
	c.Set("token", parts[1])

	return parts[1], nil
}

func (a App) GetUserFromToken(ctx *gin.Context) (user_id string, err error) {

	token, err := a.GetToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &gin.H{
			"error": "Authorization error",
		})
		log.Printf("ERROR: %v", err)
		return "", err
	}

	userDetails := map[string]string{
		"user_id":      "",
		"username":     "",
		"generated_at": "",
		"expires_at":   "",
	}

	userDetails, err = a.ValidateToken(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &gin.H{
			"error": "Authorization error",
		})
		log.Printf("ERROR: %v", err)
		return "", err
	}

	return userDetails["user_id"], nil

}
