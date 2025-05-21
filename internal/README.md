# Diretório `internal` em Aplicações Go

O diretório `internal` em uma aplicação Go é utilizado para armazenar código que deve ser acessível **apenas dentro do próprio módulo** ou projeto. Isso significa que pacotes colocados dentro de `internal` **não podem ser importados por outros módulos externos**, garantindo encapsulamento e maior segurança na arquitetura do sistema.

## Foco do Caminho `internal`

- **Restrição de acesso:** Qualquer código dentro de `internal` só pode ser utilizado pelo código que está acima dele na árvore de diretórios.
- **Organização:** É comum colocar implementações internas, helpers, serviços e lógicas que não devem ser expostas publicamente.
- **Boas práticas:** Seguir essa convenção ajuda a evitar dependências indesejadas e mantém a API pública do projeto mais limpa.

## Exemplo de Estrutura

```
/meu-projeto/
├── cmd/
├── pkg/
├── internal/
│   ├── database/
│   └── service/
└── README.md
```

Neste exemplo, apenas o código dentro de `meu-projeto` pode importar `internal/database` ou `internal/service`.

## Referências

- [Go Blog: Organizing Go code](https://blog.golang.org/organizing-go-code)
- [Documentação oficial do Go sobre internal packages](https://golang.org/doc/go1.4#internalpackages)