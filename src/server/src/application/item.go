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
		ctx.JSON(http.StatusBadRequest, &gin.H{
			"error": "Authorization error",
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

	queryString := `INSERT INTO it_wardrobe_item(
		id, 
		user_id,
		brand_name,
		garment_id,
		size_type_id,
		size_type_item_id,
		item_state
	) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	stmt, err := a.DB.Prepare(queryString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Add item error",
		})
		log.Printf("ERROR: %v", err)
		return
	}
	defer stmt.Close()
	id := uuid.New()

	randomToken := make([]byte, 32)
	_, err = rand.Read(randomToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Add item error",
		})
		log.Printf("ERROR: %v", err)
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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Add item error",
		})
		log.Printf("ERROR: %v", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"id":      id,
	})
}

func (a *App) DeleteItemHandler(ctx *gin.Context) {
	user_id, err := a.GetUserFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Authorization error",
		})
		log.Printf("ERROR: %v", err)
		return
	}

	id := ctx.Param("id")

	queryString := `UPDATE it_wardrobe_item
	SET item_state='archived'
	WHERE id=$1 AND user_id=$2;
	`
	stmt, err := a.DB.Prepare(queryString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Update error",
		})
		log.Printf("ERROR: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, user_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Delete error",
		})
		log.Printf("ERROR: %v", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success"})
}

func (a *App) GetItemHandler(ctx *gin.Context) {
	user_id, err := a.GetUserFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Authorization error",
		})
		log.Printf("ERROR: %v", err)
		return
	}

	id := ctx.Param("id")

	stmt, err := a.DB.Query(`SELECT * FROM it_wardrobe_item
	WHERE user_id=$1 and item_state=$2 and id=$3;
	`, user_id, "active", id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Get item error",
		})
		log.Printf("ERROR: %v", err)
		return
	}
	defer stmt.Close()

	var item schema.GetItem
	stmt.Next()
	if err := stmt.Scan(
		&item.Id,
		&item.UserId,
		&item.BrandName,
		&item.GarmentId,
		&item.SizeTypeId,
		&item.SizeTypeItemId,
		&item.ItemState,
	); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Get item error",
		})
		log.Printf("ERROR: %v", err)
		return
	}

	ctx.JSON(http.StatusOK, item)
}

func (a *App) ListItemsHandler(ctx *gin.Context) {
	user_id, err := a.GetUserFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Authorization error",
		})
		log.Printf("ERROR: %v", err)
		return
	}

	stmt, err := a.DB.Query(`SELECT * FROM it_wardrobe_item
	WHERE user_id=$1 and item_state=$2;
	`, user_id, "active")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Get items error",
		})
		log.Printf("ERROR: %v", err)
		return
	}
	defer stmt.Close()

	items := []schema.GetItem{}
	for stmt.Next() {
		var item schema.GetItem
		if err := stmt.Scan(
			&item.Id,
			&item.UserId,
			&item.BrandName,
			&item.GarmentId,
			&item.SizeTypeId,
			&item.SizeTypeItemId,
			&item.ItemState,
		); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Get items error",
			})
			log.Printf("ERROR: %v", err)
			return
		}
		items = append(items, item)
	}

	ctx.JSON(http.StatusOK, items)
}

func (a *App) EditItemHandler(ctx *gin.Context) {
	var req schema.AddItem
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation error",
		})
		log.Printf("ERROR: %v", err)
		return
	}

	user_id, err := a.GetUserFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Authorization error",
		})
		log.Printf("ERROR: %v", err)
		return
	}

	id := ctx.Param("id")

	stmt, err := a.DB.Query(`SELECT * FROM it_wardrobe_item
	WHERE user_id=$1 AND item_state=$2 AND id=$3;
	`, user_id, "active", id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Get item error",
		})
		log.Printf("ERROR: %v", err)
		return
	}
	defer stmt.Close()

	var item schema.GetItem
	stmt.Next()
	if err := stmt.Scan(
		&item.Id,
		&item.UserId,
		&item.BrandName,
		&item.GarmentId,
		&item.SizeTypeId,
		&item.SizeTypeItemId,
		&item.ItemState,
	); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Get item error",
		})
		log.Printf("ERROR: %v", err)
		return
	}

	stmt, err = a.DB.Query(`
	UPDATE 
		it_wardrobe_item 
	SET
		brand_name=$1,
		garment_id=$2,
		size_type_id=$3,
		size_type_item_id=$4
	WHERE id=$5`,
		req.BrandName,
		req.GarmentId,
		req.SizeTypeId,
		req.SizeTypeItemId,
		id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Update item error",
		})
		log.Printf("ERROR: %v", err)
		return
	}
	defer stmt.Close()

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"item_id": id,
	})
}
