package agent

import (
	"arboreum/internal/config"
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
	"github.com/joho/godotenv"
)

type SecondaryTwin struct {
	ctx     context.Context
	gen     *genkit.Genkit
	initErr error
}

var (
	loadEnvOnceST sync.Once
)

func NewSecondaryTwin(ctx context.Context) *SecondaryTwin {
	loadEnvOnceST.Do(loadEnvOfSecondaryTwin)

	twin := &SecondaryTwin{
		ctx: ctx,
	}

	twin.gen, twin.initErr = genkit.Init(ctx,
		genkit.WithPlugins(&googlegenai.GoogleAI{}),
		genkit.WithDefaultModel(config.LLM_SECONDARY_TWIN_NAME),
		genkit.WithPromptDir(config.DOTPROMPT_DIR),
	)

	return twin
}

func (ft *SecondaryTwin) MakeAgent() (string, error) {
	if ft.initErr != nil {
		return "", fmt.Errorf(config.PANIC_GENKIT_INIT_ERROR, ft.initErr)
	}
	prompt_file := config.DOTPROMPT_SECONDARY_TWIN_PROMPT_FILE[:len(config.DOTPROMPT_SECONDARY_TWIN_PROMPT_FILE)-len(".prompt")]
	initPrompt := genkit.LookupPrompt(ft.gen, prompt_file)
	resp, err := initPrompt.Execute(ft.ctx)
	if err != nil {
		return "", fmt.Errorf(config.PANIC_EXECUTE_PROMPT_ERROR, err)
	}
	return resp.Text(), nil
}

func loadEnvOfSecondaryTwin() {
	if err := godotenv.Load(); err != nil {
		log.Panic(config.PANIC_ENV_NOT_FOUND)
	}
	if os.Getenv("GEMINI_API_KEY") == "" {
		log.Fatal(config.PANIC_GEMINI_API_KEY_MISSING)
	}
}
