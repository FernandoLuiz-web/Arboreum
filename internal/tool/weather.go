package tool

import (
	"context"
	"fmt"

	"github.com/pontus-devoteam/agent-sdk-go/pkg/tool"
)

func NewWeatherTool() *tool.FunctionTool {
	return tool.NewFunctionTool(
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
}
