package main

import (
	"go-curso-udemy-extreme/src/repository"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	uri := "mongodb://localhost:27017"
	dbName := "curso-go"
	repoUser, errUser := repository.NewUserRepository(uri, dbName, "users")
	repoTask, errTask := repository.NewTaskRepository(uri, dbName, "tasks")

	if errUser != nil || errTask != nil {
		log.Fatalf("Erro ao iniciar repositório")
		return
	}

	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Gin!"})
	})

	app.Run(":8000")
}
