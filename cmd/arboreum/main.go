package main

import (
	"arboreum/internal/service"
	"context"
)

func main() {
	ctx := context.Background()

	service := service.NewTwinService(ctx)

	response, err := service.ProcessPrompt()
	if err != nil {
		panic(err)
	}

	println("Resposta do modelo:", response)
}
