package agent

import (
	"github.com/pontus-devoteam/agent-sdk-go/pkg/agent"
	"github.com/pontus-devoteam/agent-sdk-go/pkg/model/providers/lmstudio"

	"arboreum/internal/config"
	"arboreum/internal/tool"
)

func NewArboreumAgent() (*agent.Agent, *lmstudio.Provider) {
	provider := lmstudio.NewProvider()
	provider.SetBaseURL(config.LOCAL_SERVER_URL)
	provider.SetDefaultModel(config.LLM_NAME)

	assistant := agent.NewAgent("Arboreum")
	assistant.SetModelProvider(provider)
	assistant.WithModel(config.LLM_NAME)
	assistant.SetSystemInstructions("You are a helpful assistant.")
	assistant.WithTools(tool.NewWeatherTool())

	return assistant, provider
}
