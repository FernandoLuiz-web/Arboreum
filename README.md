# Arboreum

Sistema de IA com agentes gêmeos que refina e processa prompts usando Ollama e o modelo Gemma.

## Visão Geral

Arboreum utiliza uma arquitetura de agentes gêmeos:
- **Primeiro Gêmeo**: Refina prompts de entrada para torná-los mais eficazes
- **Segundo Gêmeo**: Processa os prompts refinados para gerar respostas

## Pré-requisitos

- Go 1.21 ou superior
- Ollama instalado e em execução
- Modelo Gemma baixado (`gemma3:latest`)

## Instalação

1. Clone o repositório:
```bash
git clone https://github.com/seunome/arboreum.git
cd arboreum
```

2. Instale as dependências:
```bash
go mod download
```

3. Certifique-se que o Ollama está rodando e o modelo Gemma está disponível:
```bash
ollama list
# Se gemma3 não estiver listado, execute:
ollama pull gemma3
```

## Configuração

1. Crie um arquivo `.env` na raiz do projeto:
```env
GEMINI_API_KEY=sua_chave_api_aqui
```

2. A aplicação criará um diretório `prompt` para armazenar resultados intermediários.

## Como Usar

Execute a aplicação:
```bash
go run cmd/arboreum/main.go
```

O sistema irá:
1. Receber um prompt de entrada
2. Refiná-lo usando o Primeiro Gêmeo
3. Salvar o prompt refinado em um arquivo
4. Processá-lo usando o Segundo Gêmeo
5. Exibir o resultado final

## Estrutura do Projeto

```
.
├── cmd/
│   └── arboreum/
│       └── main.go
├── internal/
│   ├── agent/
│   │   ├── primary_twin.go
│   │   └── secondary_twin.go
│   ├── config/
│   └── model/
├── prompt/
└── .env
```

## Tratamento de Erros

A aplicação inclui tratamento abrangente de erros para:
- Problemas de conexão com Ollama
- Problemas de carregamento do modelo
- Operações com arquivos
- Timeouts de contexto

### Notas de Uso

- Certifique-se de que o Ollama esteja rodando antes de executar a aplicação
- O diretório `prompt` será criado automaticamente na primeira execução
- Os logs de erro são detalhados para facilitar a depuração

### Requisitos de Sistema

- Windows 10/11 ou Linux/MacOS
- Mínimo de 8GB de RAM recomendado
- Espaço em disco suficiente para os modelos do Ollama