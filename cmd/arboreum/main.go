package main

import (
	"context"
	"log"
	"os"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Obtém a chave da API do Gemini do ambiente
	geminiAPIKey := os.Getenv("GEMINI_API_KEY")
	if geminiAPIKey == "" {
		log.Fatal("GEMINI_API_KEY não definida no ambiente")
	}

	// Inicializa o Genkit com o plugin Google AI e o Gemini 2.0 Flash.
	g, err := genkit.Init(ctx,
		genkit.WithPlugins(&googlegenai.GoogleAI{}),
		genkit.WithDefaultModel("googleai/gemini-2.0-flash"),
	)
	if err != nil {
		log.Fatalf("Não foi possível inicializar o Genkit: %v", err)
	}

	resp, err := genkit.Generate(ctx, g, ai.WithPrompt("Qual é o sentido da vida?"))
	if err != nil {
		log.Fatalf("Não foi possível gerar a resposta do modelo: %v", err)
	}

	log.Println(resp.Text())
}
