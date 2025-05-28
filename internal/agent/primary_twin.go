package agent

import (
	"arboreum/internal/config"
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
	ctx     context.Context
	gen     *genkit.Genkit
	model   ai.Model
	initErr error
}

var (
	loadEnvOnceFT sync.Once
)

func NewFirstTwin(ctx context.Context) *FirstTwin {
	loadEnvOnceFT.Do(loadEnvOfFirstTwin)

	ollamaPlugin := &ollama.Ollama{ServerAddress: config.LOCALHOST_OLLAMA_SERVER}

	twin := &FirstTwin{
		ctx: ctx,
	}
	twin.gen, twin.initErr = genkit.Init(ctx,
		genkit.WithPlugins(ollamaPlugin),
		genkit.WithPromptDir("prompt"),
	)

	twin.model = model.DefineFirstTwinModel(ollamaPlugin, twin.gen)

	return twin
}

func (ft *FirstTwin) RefinePrompt() (string, error) {
	if ft.initErr != nil {
		return "", fmt.Errorf("erro na inicialização do Genkit: %w", ft.initErr)
	}

	initPrompt := genkit.LookupPrompt(ft.gen, "primary_twin")
	// resp, err := genkit.Generate(ft.ctx, ft.gen,
	// 	ai.WithModel(ft.model),
	// )

	resp, err := initPrompt.Execute(ft.ctx, ai.WithModel(ft.model))

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
