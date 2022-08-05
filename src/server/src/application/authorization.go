package application

import (
	"log"
	"net/http"

	schema "main/src/schemas"

	"github.com/gin-gonic/gin"
)

func (a *App) CreateAccount(ctx *gin.Context) {
	var req schema.CreateAccountRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation error",
		})
		return
	}

	response, err := a.RegisterUser(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Registration error",
		})
		log.Printf("ERROR: %v", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": response})
}

func (a *App) AuthenticationsHandler(ctx *gin.Context) {
	var req schema.AuthenticationRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation error",
		})
		log.Printf("ERROR: %v", err)
		return
	}

	tokenDetails, err := a.GenerateToken(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Generate Token error",
		})
		log.Printf("ERROR: %v", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenDetails})
}
