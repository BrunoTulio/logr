# 🚀 Golr - Biblioteca de Logging Unificada para Go

[![Go Version](https://img.shields.io/badge/go-1.22.10+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/BrunoTulio/logr)](https://goreportcard.com/report/github.com/BrunoTulio/logr)

Uma biblioteca de logging moderna e flexível para Go que oferece uma interface unificada para múltiplas implementações de logging, permitindo trocar entre diferentes backends sem alterar seu código.

## ✨ Características

- 🎯 **Interface Unificada**: Uma única API para diferentes implementações de logging
- 🔄 **Múltiplos Backends**: Suporte para atualmente `slog` (padrão Go) e `zap` (alta performance)
- 🌐 **Logging Global**: Sistema de logging global opcional com funções de conveniência
- 📝 **Campos Estruturados**: Sistema robusto de campos tipados
- 🔗 **Contexto**: Propagação de campos via context
- 📊 **Múltiplos Outputs**: Console e arquivo simultaneamente
- 🔄 **Rotação de Logs**: Rotação automática com lumberjack
- 🎨 **Formatação Flexível**: JSON e TEXT
- 🛡️ **Type Safety**: Interface bem definida com validação de tipos

## 🚀 Instalação

```bash
go get github.com/BrunoTulio/logr
```

## 📖 Uso Básico

### Logger Direto (Recomendado)

```go
package main

import (
    "github.com/BrunoTulio/logr"
    "github.com/BrunoTulio/logr/adapters/slog.v1"
)

func main() {
    // Criar logger com slog (padrão do Go)
    logger := slog.New(
        slog.WithConsole(true),
        slog.WithConsoleLevel("INFO"),
        slog.WithConsoleFormatter("TEXT"),
    )
    
    logger.Info("Aplicação iniciada")
    logger.Error("Erro encontrado")
    
    // Com campos estruturados
    logger.WithFields(
        logr.String("user_id", "123"),
        logr.Bool("active", true),
    ).Info("Usuário logado")
}
```

### Logger com Zap (Alta Performance)

```go
package main

import (
    "github.com/BrunoTulio/logr"
    "github.com/BrunoTulio/logr/adapters/zap.v1"
)

func main() {
    // Criar logger com zap para alta performance
    logger := zap.New(
        zap.WithConsole(true),
        zap.WithConsoleLevel("INFO"),
        zap.WithConsoleFormatter("JSON"),
        zap.WithFile(true, "/var/log", "app.log"),
        zap.WithFileRotation(100, 30, true), // 100MB, 30 dias, comprimir
    )
    
    logger.Info("Sistema iniciado")
    
    // Com campos agrupados
    userLogger := logger.WithFields(
        logr.Group("user", 
            logr.String("name", "João"),
            logr.Bool("active", true),
        ),
        logr.Int("age", 30),
    )
    
    userLogger.Info("Usuário processado")
}
```

### Logger Global (Opcional)

```go
package main

import (
    "github.com/BrunoTulio/logr"
    "github.com/BrunoTulio/logr/adapters/slog.v1"
)

func main() {
    // Configurar o logger global (apenas se necessário)
    logger := slog.New(
        slog.WithConsole(true),
        slog.WithConsoleFormatter("JSON"),
    )
    
    logr.Set(logger)
    
    // Usar funções globais
    logr.Info("Usando logger global")
    logr.WithFields(
        logr.String("component", "global"),
    ).Info("Sistema inicializado")
}
```

### Usando Contexto

```go
package main

import (
    "context"
    "github.com/BrunoTulio/logr"
    "github.com/BrunoTulio/logr/adapters/slog.v1"
)

func processUser(ctx context.Context, userID string) {
    logger := slog.New(slog.WithConsole(true))
    
    // Adicionar campos ao contexto
    ctx = logger.WithFields(
        logr.String("request_id", "req-12345"),
        logr.String("user_id", userID),
    ).ToContext(ctx)
    
    // Em outra função
    processRequest(ctx)
}

func processRequest(ctx context.Context) {
    logger := slog.New(slog.WithConsole(true))
    requestLogger := logger.FromContext(ctx)
    
    requestLogger.Info("Processando requisição")
}
```

## 🔧 Configuração Avançada

### Configuração Completa

```go
// Com slog
logger := slog.New(
    slog.WithConsole(true),
    slog.WithConsoleLevel("DEBUG"),
    slog.WithConsoleFormatter("TEXT"),
    slog.WithAddSource(true),
    
    slog.WithFile(true, "/var/log", "application.log"),
    slog.WithFileLevel("INFO"),
    slog.WithFileFormatter("JSON"),
    slog.WithFileRotation(100, 30, true),
)

// Com zap
logger := zap.New(
    zap.WithConsole(true),
    zap.WithConsoleLevel("INFO"),
    zap.WithConsoleFormatter("JSON"),
    
    zap.WithFile(true, "/var/log", "application.log"),
    zap.WithFileLevel("DEBUG"),
    zap.WithFileFormatter("JSON"),
    zap.WithFileRotation(100, 30, true),
)
```

## 📊 Tipos de Campos Suportados

```go
logger.WithFields(
    logr.String("name", "João"),
    logr.Bool("active", true),
    logr.Int("age", 30),
    logr.Uint64("id", 123456789),
    logr.Float64("score", 95.5),
    logr.Time("created_at", time.Now()),
    logr.Duration("duration", time.Second*5),
    logr.Group("address",
        logr.String("street", "Rua das Flores"),
        logr.String("city", "São Paulo"),
    ),
).Info("Dados do usuário")
```

## 🏗️ Arquitetura

```
golr/
├── logger.go          # Interface principal
├── level.go           # Níveis de log
├── field.go           # Sistema de campos
├── global.go          # Logger global (opcional)
├── noop.go           # Implementação vazia
└── adapters/
    ├── slog.v1/       # Implementação com slog (padrão Go)
    └── zap.v1/        # Implementação com zap (alta performance)
```

## 🎯 Vantagens de Usar golr

### 1. **Flexibilidade**

- Troque entre implementações sem alterar o código
- Configure diferentes outputs (console/arquivo) independentemente
- Suporte a múltiplos formatos (JSON/TEXT)

### 2. **Simplicidade**

- API limpa e intuitiva
- Logger global opcional para casos específicos
- Integração com context para propagação de campos

### 3. **Performance**

- Escolha entre slog (padrão) e zap (alta performance)
- Campos estruturados eficientes
- Rotação de logs automática

### 4. **Manutenibilidade**

- Interface bem definida
- Código limpo e testável
- Fácil de estender com novos adapters

### 5. **Compatibilidade**

- Suporte ao slog padrão do Go 1.21+
- Compatível com ferramentas de logging existentes
- Migração fácil de outras bibliotecas

## 🔮 Implementações Futuras

### 🧪 Testes e Qualidade

- [ ] **Testes Unitários**: Cobertura completa de todos os adapters
- [ ] **Testes de Integração**: Testes end-to-end com diferentes configurações
- [ ] **Benchmarks**: Comparação de performance entre adapters
- [ ] **CI/CD**: Pipeline automatizado com GitHub Actions

### 📚 Documentação e Exemplos

- [ ] **Documentação da API**: Godoc completo com exemplos
- [ ] **Best Practices**: Padrões recomendados de uso
- [ ] **Tutorial Interativo**: Exemplos práticos passo a passo

### 🔌 Adapters Disponíveis

- ✅ **Slog** – Adapter para [log/slog](https://pkg.go.dev/log/slog) (Go 1.21+)
- ✅ **Zap** – Adapter para [uber-go/zap](https://github.com/uber-go/zap)

### 🔌 Novos Adapters

- [ ] **Zerolog**: Adapter para [zerolog](https://github.com/rs/zerolog) - JSON estruturado
- [ ] **Logrus**: Adapter para [logrus](https://github.com/sirupsen/logrus) - Compatibilidade

### 🚀 Funcionalidades Avançadas

- [ ] **Sampling**: Para logs de alta frequência
- [ ] **Métricas**: Integração com Prometheus/OpenTelemetry
- [ ] **Buffering**: Buffer configurável para melhor performance
- [x] **Compressão**: Compressão automática de logs antigos
- [ ] **Middleware**: Intercepta o log antes de ser enviado ao destino final

## 🤝 Contribuindo

Contribuições são muito bem-vindas! Por favor:

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

### Tipos de Contribuição

- 🐛 **Bug Fixes**: Correção de bugs
- ✨ **New Features**: Novas funcionalidades
- 📚 **Documentation**: Melhorias na documentação
- 🧪 **Tests**: Adição de testes
- 🔌 **New Adapters**: Novos adapters para bibliotecas de logging
- 🎨 **UI/UX**: Melhorias na experiência do usuário

## 📄 Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

## 🙏 Agradecimentos

- [Go Slog](https://pkg.go.dev/log/slog) - Pela interface padrão do Go
- [Uber Zap](https://github.com/uber-go/zap) - Por inspirar a alta performance
- [Lumberjack](https://github.com/natefinch/lumberjack) - Pela rotação de logs
- [Zerolog](https://github.com/rs/zerolog) - Por inspirar o JSON estruturado
- [Logrus](https://github.com/sirupsen/logrus) - Por popularizar o logging estruturado

## 📞 Contato

- **Autor**: Bruno Tulio
- **GitHub**: [@BrunoTulio](https://github.com/BrunoTulio)

---

⭐ **Se este projeto foi útil, considere dar uma estrela!** ⭐
