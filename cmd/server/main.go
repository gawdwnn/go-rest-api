package main

import (
	"fmt"
	"net/http"

	"github.com/Smarttin/go-rest-api/internal/comment"
	"github.com/Smarttin/go-rest-api/internal/database"
	transportHTTP "github.com/Smarttin/go-rest-api/internal/transport/http"
)

// struct which contains things like pointers to DB connections
type App struct{}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our app")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set  up server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("go rest api")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting api")
		fmt.Println(err)
	}
}
