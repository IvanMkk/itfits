package application

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) OptionsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Access-Control-Allow-Methods": "POST, GET, OPTIONS",
		"Access-Control-Allow-Headers": "Content-Type",
	})
}
