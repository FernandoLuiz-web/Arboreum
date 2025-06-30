package service

import (
	"arboreum/internal/agent"
	"arboreum/internal/config"
	"context"
	"fmt"
	"os"
	"path/filepath"
)

type TwinService struct {
	ctx context.Context
}

func NewTwinService(ctx context.Context) *TwinService {
	return &TwinService{ctx: ctx}
}

func (s *TwinService) ProcessPrompt() (string, error) {
	refinedPrompt, err := s.executePrimaryTwin()
	if err != nil {
		return "", fmt.Errorf(config.PANIC_REFINE_PROMPT_ERROR, err)
	}

	if err := s.savePrompt(refinedPrompt); err != nil {
		return "", fmt.Errorf(config.PANIC_WRITE_PROMPT_ERROR, err)
	}

	return s.executeSecondaryTwin()
}

func (s *TwinService) executePrimaryTwin() (string, error) {
	twin := agent.NewPrimaryTwin(s.ctx)
	return twin.RefinePrompt()
}

func (s *TwinService) savePrompt(prompt string) error {
	promptDir := config.DOTPROMPT_DIR
	promptFileName := config.DOTPROMPT_SECONDARY_TWIN_PROMPT_FILE

	if err := os.MkdirAll(promptDir, 0755); err != nil {
		return err
	}

	promptFilePath := filepath.Join(promptDir, promptFileName)
	return os.WriteFile(promptFilePath, []byte(prompt), 0644)
}

func (s *TwinService) executeSecondaryTwin() (string, error) {
	twin := agent.NewSecondaryTwin(s.ctx)
	return twin.MakeAgent()
}
