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

	return &App{
		Port:    os.Getenv("PORT"),
		DB:      db,
		Methods: gin.Default(),
	}
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

	a.Methods.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	a.Methods.POST("/v1/registration", a.CreateAccount)
	a.Methods.POST("/v1/authentications", a.AuthenticationsHandler)

	a.Methods.Run(fmt.Sprintf("0.0.0.0:%v", a.Port))
}
