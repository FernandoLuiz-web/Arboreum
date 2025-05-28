package model

import (
	"arboreum/internal/config"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/ollama"
)

func DefineFirstTwinModel(o *ollama.Ollama, g *genkit.Genkit) ai.Model {
	return o.DefineModel(g,
		ollama.ModelDefinition{
			Name: config.LLM_PRIMARY_TWIN_NAME,
			Type: "chat",
		},
		&ai.ModelInfo{
			Supports: &ai.ModelSupports{
				Multiturn:  true,  // Permite conversas com múltiplas interações (multi-turn dialogue).
				SystemRole: true,  // Suporta o uso de mensagens de sistema (system role messages).
				Tools:      false, // Não suporta integração com ferramentas externas (tools).
				Media:      false, // Não suporta entrada ou saída de mídia (media input/output).
			},
		},
	)
}

func DefineMessageAi(prompt string) *ai.Message {
	return &ai.Message{
		Role:    "user",
		Content: []*ai.Part{{Text: prompt}},
	}
}
