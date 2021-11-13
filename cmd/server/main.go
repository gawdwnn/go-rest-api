package main

import "fmt"

// struct which contains things like pointers to DB connections
type App struct {}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our app")
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
