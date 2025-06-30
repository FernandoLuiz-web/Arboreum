package config

const (
	// Panic Messages
	PANIC_ENV_NOT_FOUND          = "Aviso: .env não encontrado."
	PANIC_GEMINI_API_KEY_MISSING = "GEMINI_API_KEY não definida no ambiente"
	PANIC_GENKIT_INIT_ERROR      = "erro na inicialização do Genkit: %w"
	PANIC_EXECUTE_PROMPT_ERROR   = "erro ao executar o prompt: %w"
	PANIC_REFINE_PROMPT_ERROR    = "erro ao refinar prompt: %w"
	PANIC_WRITE_PROMPT_ERROR     = "erro ao escrever o prompt refinado: %w"
	PANIC_SECONDARY_TWIN_ERROR   = "erro ao executar segundo gêmeo: %w"

	// Warning Messages
	WARN_MODEL_NOT_LOADED = "Aviso: Modelo %s não carregado completamente"
	WARN_SLOW_RESPONSE    = "Aviso: Resposta do modelo está demorando mais que o esperado"
	WARN_DIRECTORY_CREATE = "Aviso: Criando diretório %s"

	// Info Messages
	INFO_MODEL_LOADED        = "Modelo definido: %s"
	INFO_PROCESSING_PROMPT   = "Processando prompt com %s"
	INFO_GENERATING_RESPONSE = "Gerando resposta..."
	INFO_SAVING_PROMPT       = "Salvando prompt refinado em: %s"
)
