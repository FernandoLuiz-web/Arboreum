package main

import (
	"context"
	"fmt"
	"log"

	"github.com/pontus-devoteam/agent-sdk-go/pkg/agent"
	"github.com/pontus-devoteam/agent-sdk-go/pkg/model/providers/lmstudio"
	"github.com/pontus-devoteam/agent-sdk-go/pkg/runner"
	"github.com/pontus-devoteam/agent-sdk-go/pkg/tool"
)

const (
	LOCAL_SERVER_URL = "http://127.0.0.1:1234/v1"
	LLM_NAME         = "openhermes-2.5-mistral-7b"
)

func main() {
	provider := lmstudio.NewProvider()
	provider.SetBaseURL(LOCAL_SERVER_URL)
	provider.SetDefaultModel(LLM_NAME)
	assistant := agent.NewAgent("Assistant")
	assistant.SetModelProvider(provider)
	assistant.WithModel(LLM_NAME)
	assistant.SetSystemInstructions("You are a helpful assistant.")

	tool := tool.NewFunctionTool(
		"get_weather",
		"Get the weather for a city",
		func(ctx context.Context, params map[string]interface{}) (interface{}, error) {
			city := params["city"].(string)
			return fmt.Sprintf("The weather in %s is sunny.", city), nil
		},
	).WithSchema(map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"city": map[string]interface{}{
				"type":        "string",
				"description": "The city to get weather for",
			},
		},
		"required": []string{"city"},
	})

	assistant.WithTools(tool)

	runned := runner.NewRunner()
	runned.WithDefaultProvider(provider)

	result, err := runned.RunSync(assistant, &runner.RunOptions{
		Input: "What's the weather in Tokyo?",
	})
	if err != nil {
		log.Fatalf("Error running agent: %v", err)
	}

	fmt.Println(result.FinalOutput)
}
