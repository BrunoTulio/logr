# 📚 Exemplos de Uso - Golog

Este diretório contém exemplos práticos de como usar a biblioteca Golog.

## 🚀 Como Executar os Exemplos

### 1. Exemplo Básico
Demonstra o uso básico com slog e zap:

```bash
cd basic
go run simple_example.go
```

### 2. Exemplo Completo
Demonstra todas as funcionalidades da biblioteca:

```bash
cd complete
go run complete_example.go
```

### 3. Exemplos de Configuração
Mostra diferentes configurações para vários ambientes:

```bash
cd config
go run config_examples.go
```

## 📁 Estrutura dos Exemplos

```
examples/
├── basic/
│   ├── simple_example.go    # Exemplo básico
│   └── go.mod
├── complete/
│   ├── complete_example.go  # Exemplo completo
│   └── go.mod
├── config/
│   ├── config_examples.go   # Configurações
│   └── go.mod
└── README.md
```

## 🎯 O que Cada Exemplo Demonstra

### Basic Example
- Logger básico com slog e zap
- Campos estruturados simples
- Logger com arquivo

### Complete Example
- Todas as funcionalidades da biblioteca
- Uso com contexto
- Logger global
- Campos agrupados

### Config Examples
- Configurações para desenvolvimento
- Configurações para produção
- Configurações para Docker
- Configurações para microserviços
- Configurações para alta performance
- Configurações para debugging

## 🔧 Pré-requisitos

- Go 1.22.10 ou superior
- A biblioteca Golog instalada

## 📝 Notas

- Cada exemplo tem seu próprio `go.mod` para evitar conflitos
- Os exemplos usam `replace` para referenciar a biblioteca local
- Execute `go mod tidy` se houver problemas de dependências
