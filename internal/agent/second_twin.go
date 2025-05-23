package agent

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
	"github.com/joho/godotenv"
)

type SecondTwin struct {
	ctx     context.Context
	gen     *genkit.Genkit
	initErr error
}

var (
	loadEnvOnceST sync.Once
)

func NewSecondTwin(ctx context.Context) *SecondTwin {
	loadEnvOnceST.Do(loadEnvOfSecondTwin)

	twin := &SecondTwin{
		ctx: ctx,
	}

	twin.gen, twin.initErr = genkit.Init(ctx,
		genkit.WithPlugins(&googlegenai.GoogleAI{}),
		genkit.WithDefaultModel("googleai/gemini-2.0-flash"),
		genkit.WithPromptDir("prompt"),
	)

	return twin
}

func (ft *SecondTwin) MakeAgent() (string, error) {
	if ft.initErr != nil {
		return "", fmt.Errorf("erro na inicialização do Genkit: %w", ft.initErr)
	}

	initPrompt := genkit.LookupPrompt(ft.gen, "second_twin")
	resp, err := initPrompt.Execute(ft.ctx)
	if err != nil {
		return "", fmt.Errorf("erro ao executar o prompt: %w", err)
	}
	return resp.Text(), nil
}

func loadEnvOfSecondTwin() {
	if err := godotenv.Load(); err != nil {
		log.Panic("Aviso: .env não encontrado.")
	}
	if os.Getenv("GEMINI_API_KEY") == "" {
		log.Fatal("GEMINI_API_KEY não definida no ambiente")
	}
}
