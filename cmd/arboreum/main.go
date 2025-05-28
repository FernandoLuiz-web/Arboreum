package main

import (
	"arboreum/internal/agent"
	"context"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	ctx := context.Background()

	refinedPrompt, err := runFirstTwin(ctx)
	if err != nil {
		panic(fmt.Errorf("erro ao refinar prompt: %w", err))
	}

	err = writePromptToFile(refinedPrompt)
	if err != nil {
		panic(fmt.Errorf("erro ao escrever o prompt refinado: %w", err))
	}

	result, err := runSecondTwin(ctx)
	if err != nil {
		panic(fmt.Errorf("erro ao executar segundo gêmeo: %w", err))
	}

	fmt.Println("\nResposta do segundo gêmeo:")
	fmt.Println(result)
}

func runFirstTwin(ctx context.Context) (string, error) {
	twin := agent.NewFirstTwin(ctx)
	return twin.RefinePrompt()
}

func writePromptToFile(prompt string) error {
	promptDir := "prompt"
	promptFileName := "secondary_twin.prompt"

	if err := os.MkdirAll(promptDir, 0755); err != nil {
		return err
	}

	promptFilePath := filepath.Join(promptDir, promptFileName)

	return os.WriteFile(promptFilePath, []byte(prompt), 0644)
}

func runSecondTwin(ctx context.Context) (string, error) {
	twin := agent.NewSecondTwin(ctx)
	return twin.MakeAgent()
}
