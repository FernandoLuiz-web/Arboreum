package agent

import (
	"arboreum/internal/model"
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/ollama"
	"github.com/joho/godotenv"
)

type FirstTwin struct {
	prompt  string
	ctx     context.Context
	gen     *genkit.Genkit
	initErr error
}

var (
	loadEnvOnceFT  sync.Once
	firstTwinModel ai.Model
)

func NewFirstTwin(ctx context.Context, prompt string) *FirstTwin {
	loadEnvOnceFT.Do(loadEnvOfFirstTwin)
	ollamaPlugin := &ollama.Ollama{ServerAddress: "http://127.0.0.1:11434"}

	twin := &FirstTwin{
		prompt: prompt,
		ctx:    ctx,
	}

	twin.gen, twin.initErr = genkit.Init(ctx,
		genkit.WithPlugins(ollamaPlugin))

	firstTwinModel = model.DefineFirstTwinModel(ollamaPlugin, twin.gen)

	return twin
}

func (ft *FirstTwin) RefinePrompt() (string, error) {
	if ft.initErr != nil {
		return "", fmt.Errorf("erro na inicialização do Genkit: %w", ft.initErr)
	}

	resp, err := genkit.Generate(ft.ctx, ft.gen, ai.WithModel(firstTwinModel), ai.WithPrompt(ft.prompt))
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
