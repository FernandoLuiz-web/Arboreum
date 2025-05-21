## Função do diretório `pkg` em aplicações Go

No ecossistema Go, o diretório `pkg` é uma convenção utilizada para armazenar código que pode ser reutilizado por outros projetos ou módulos. Ele geralmente contém bibliotecas, utilitários ou componentes que não dependem diretamente da aplicação principal, mas que podem ser importados por ela ou por outros projetos Go.

### Vantagens de usar o diretório `pkg`:

- **Organização:** Facilita a separação entre o código reutilizável (em `pkg`) e o código específico da aplicação (geralmente em `cmd` ou `internal`).
- **Reutilização:** Permite que outros projetos importem facilmente os pacotes do diretório `pkg`.
- **Manutenção:** Torna o código mais modular e fácil de manter.

### Estrutura típica:

```
meu-projeto/
├── cmd/        # Código da aplicação principal (entrypoints)
├── internal/   # Código privado, não exportado para outros projetos
├── pkg/        # Pacotes reutilizáveis e exportáveis
└── go.mod
```

> **Observação:** O uso do diretório `pkg` é uma convenção popular, mas não obrigatória. O importante é manter o código organizado e fácil de entender para outros desenvolvedores.