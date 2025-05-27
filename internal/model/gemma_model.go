package model

import (
	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/ollama"
)

func DefineFirstTwinModel(o *ollama.Ollama, g *genkit.Genkit) ai.Model {
	return o.DefineModel(g,
		ollama.ModelDefinition{
			Name: "gemma3",
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
