package main

import (
	"arboreum/internal/agent"
	"context"
	"fmt"
	"os"
	"path/filepath"
)

const (
	PromptInputText = `Você é um assistente de IA que ajuda a refinar prompts para o Gemini AI. Por favor refine o seguinte prompt para torná-lo mais claro e eficaz: preciso de um secretário.`
	PromptDir       = "prompt"
	PromptFileName  = "second_twin.prompt"
)

func main() {
	ctx := context.Background()

	refinedPrompt, err := runFirstTwin(ctx, PromptInputText)
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

func runFirstTwin(ctx context.Context, input string) (string, error) {
	twin := agent.NewFirstTwin(ctx, input)
	return twin.RefinePrompt()
}

func writePromptToFile(prompt string) error {
	if err := os.MkdirAll(PromptDir, 0755); err != nil {
		return err
	}

	promptFilePath := filepath.Join(PromptDir, PromptFileName)

	return os.WriteFile(promptFilePath, []byte(prompt), 0644)
}

func runSecondTwin(ctx context.Context) (string, error) {
	twin := agent.NewSecondTwin(ctx)
	return twin.MakeAgent()
}
