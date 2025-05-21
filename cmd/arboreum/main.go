package main

import (
	"fmt"
	"log"

	"arboreum/internal/agent"

	"github.com/pontus-devoteam/agent-sdk-go/pkg/runner"
)

func main() {
	assistant, provider := agent.NewArboreumAgent()

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
