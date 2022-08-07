package application

import (
	"crypto/rand"
	"log"
	schema "main/src/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (a *App) AddItemHandler(ctx *gin.Context) {
	user_id, err := a.GetUserFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation error",
		})
		return
	}

	var req schema.AddItem
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation error",
		})
		log.Printf("ERROR: %v", err)
		return
	}

	queryString := `insert into it_users(
		id, 
		user_id
		brand_name,
		garment_id,
		size_type_id,
		size_type_item_id,
		item_state
	) values ($1, $2, $3, $4, $5, $6, $7)`
	stmt, err := a.DB.Prepare(queryString)
	if err != nil {
		return
	}
	defer stmt.Close()
	id := uuid.New()

	randomToken := make([]byte, 32)
	_, err = rand.Read(randomToken)
	if err != nil {
		return
	}

	_, err = stmt.Exec(id,
		user_id,
		req.BrandName,
		req.GarmentId,
		req.SizeTypeId,
		req.SizeTypeItemId,
		"active")

	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success"})
}
