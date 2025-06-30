package main

import (
	"arboreum/internal/service"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// twin comunication function
func twin() {
	ctx := context.Background()
	service := service.NewTwinService(ctx)
	response, err := service.ProcessPrompt()
	if err != nil {
		panic(err)
	}
	println("Resposta do modelo:", response)
}
