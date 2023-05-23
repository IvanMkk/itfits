package application

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	timeLayout = "2006-01-02 15:04:05"
	timeLayoutPQ = "2006-01-02T15:04:05Z"
)

type App struct {
	Port    string
	DB      *sql.DB
	Methods *gin.Engine
}

type Client struct {
}

func New() *App {
	readEnv()

	db, err := sql.Open("postgres", os.Getenv("POSTGRES"))
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		db.Close()
		log.Fatalf("ERROR: DB connection lost. %s", err)
	}

	app := &App{
		Port:    os.Getenv("PORT"),
		DB:      db,
		Methods: gin.Default(),
	}

	app.Methods.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Type")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Next()
	})

	app.Methods.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	app.Methods.POST("/v1/users", app.RegistrationHandler)
	app.Methods.GET("/v1/users", app.AuthenticationsHandler)
	app.Methods.OPTIONS("/v1/users", app.OptionsHandler)

	app.Methods.POST("/v1/item", app.AddItemHandler)
	app.Methods.DELETE("/v1/item/:id", app.DeleteItemHandler)
	app.Methods.GET("/v1/item", app.ListItemsHandler)
	app.Methods.GET("/v1/item/:id", app.GetItemHandler)
	app.Methods.PUT("/v1/item/:id", app.EditItemHandler)

	return app
}

func readEnv() {
	if os.Getenv("PORT") == "" {
		file, err := os.Open("test.env")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		line := bufio.NewScanner(file)
		for line.Scan() {
			fmt.Println(line.Text())
			s := strings.Split(line.Text(), ":")
			os.Setenv(s[0], strings.Join(s[1:], ":"))
		}

		if err := line.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

func (a *App) Start() {
	a.Methods.Use(gin.Logger())

	a.Methods.Run(fmt.Sprintf("0.0.0.0:%v", a.Port))
}
