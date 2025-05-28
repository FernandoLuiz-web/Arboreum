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

type PrimaryTwin struct {
	ctx     context.Context
	gen     *genkit.Genkit
	model   ai.Model
	initErr error
}

var (
	loadEnvOnceFT sync.Once
)

func NewPrimaryTwin(ctx context.Context) *PrimaryTwin {
	loadEnvOnceFT.Do(loadEnvOfPrimaryTwin)

	ollamaPlugin := &ollama.Ollama{ServerAddress: config.LOCALHOST_OLLAMA_SERVER}

	twin := &PrimaryTwin{
		ctx: ctx,
	}
	twin.gen, twin.initErr = genkit.Init(ctx,
		genkit.WithPlugins(ollamaPlugin),
		genkit.WithPromptDir(config.DOTPROMPT_DIR),
	)

	twin.model = model.DefineFirstTwinModel(ollamaPlugin, twin.gen)

	return twin
}

func (ft *PrimaryTwin) RefinePrompt() (string, error) {
	if ft.initErr != nil {
		return "", fmt.Errorf(config.PANIC_GENKIT_INIT_ERROR, ft.initErr)
	}
	prompt_file := config.DOTPROMPT_PRIMARY_TWIN_PROMPT_FILE[:len(config.DOTPROMPT_PRIMARY_TWIN_PROMPT_FILE)-len(".prompt")]
	initPrompt := genkit.LookupPrompt(ft.gen, prompt_file)
	resp, err := initPrompt.Execute(ft.ctx, ai.WithModel(ft.model))
	if err != nil {
		return "", fmt.Errorf(config.PANIC_EXECUTE_PROMPT_ERROR, err)
	}
	return resp.Text(), nil
}

func loadEnvOfPrimaryTwin() {
	if err := godotenv.Load(); err != nil {
		log.Panic(config.PANIC_ENV_NOT_FOUND)
	}
	if os.Getenv("GEMINI_API_KEY") == "" {
		log.Fatal(config.PANIC_GEMINI_API_KEY_MISSING)
	}
}
