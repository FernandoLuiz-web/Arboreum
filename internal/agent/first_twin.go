package agent

import (
	"arboreum/internal/config"
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
	"github.com/joho/godotenv"
)

type FirstTwin struct {
	prompt  string
	ctx     context.Context
	gen     *genkit.Genkit
	initErr error
}

var (
	loadEnvOnceFT sync.Once
)

func NewFirstTwin(ctx context.Context, prompt string) *FirstTwin {
	loadEnvOnceFT.Do(loadEnvOfFirstTwin)

	twin := &FirstTwin{
		prompt: prompt,
		ctx:    ctx,
	}

	twin.gen, twin.initErr = genkit.Init(ctx,
		genkit.WithPlugins(&googlegenai.GoogleAI{}),
		genkit.WithDefaultModel(config.LLM_NAME),
	)

	return twin
}

func (ft *FirstTwin) RefinePrompt() (string, error) {
	if ft.initErr != nil {
		return "", fmt.Errorf("erro na inicialização do Genkit: %w", ft.initErr)
	}

	resp, err := genkit.Generate(ft.ctx, ft.gen, ai.WithPrompt(ft.prompt))
	if err != nil {
		return "", fmt.Errorf("erro ao gerar resposta da IA: %w", err)
	}

	return resp.Text(), nil
}

func loadEnvOfFirstTwin() {
	if err := godotenv.Load(); err != nil {
		log.Panic("Aviso: .env não encontrado.")
	}
	if os.Getenv("GEMINI_API_KEY") == "" {
		log.Fatal("GEMINI_API_KEY não definida no ambiente")
	}
}
