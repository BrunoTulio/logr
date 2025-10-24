package main

import (
	"context"
	"time"

	"github.com/BrunoTulio/logr"
	"github.com/BrunoTulio/logr/adapters/slog.v1"
	"github.com/BrunoTulio/logr/adapters/zap.v1"
)

const (
	maxFileSizeMB      = 10 // tamanho máximo do arquivo em MB
	retentionDays      = 7  // quantidade de dias para manter os logs
	loginCount         = 42
	scoreValue         = 95.7
	sessionHours       = 2
	totalAmount        = 299.99
	sleepDuration      = 100 * time.Millisecond
	processingDuration = 100 * time.Millisecond
)

func main() {
	// Exemplo 1: Logger com Slog (padrão do Go)
	demonstrateSlog()

	// Exemplo 2: Logger com Zap (alta performance)
	demonstrateZap()

	// Exemplo 3: Logger Global
	demonstrateGlobalLogger()

	// Exemplo 4: Uso com Contexto
	demonstrateContextUsage()
}

func demonstrateSlog() {
	// Logger básico com console
	logger := slog.New(
		slog.WithConsole(true),
		slog.WithConsoleLevel("DEBUG"),
		slog.WithConsoleFormatter("TEXT"),
		slog.WithAddSource(true),
	)

	logger.Info("=== Demonstração Slog ===")
	logger.Info("Logger criado com slog")
	logger.Debug("Mensagem de debug")
	logger.Warn("Aviso importante")
	logger.Error("Erro encontrado")

	// Logger com arquivo
	fileLogger := slog.New(
		slog.WithConsole(true),
		slog.WithFile(true, "./logs", "slog_example.log"),
		slog.WithFileFormatter("JSON"),
		slog.WithFileRotation(maxFileSizeMB, retentionDays, true), // 10MB, 7 dias, comprimir
	)

	fileLogger.WithFields(
		logr.String("service", "example"),
		logr.Int("version", 1),
	).Info("Log salvo em arquivo")
}

func demonstrateZap() {
	// Logger de alta performance
	logger := zap.New(
		zap.WithConsole(true),
		zap.WithConsoleLevel("INFO"),
		zap.WithConsoleFormatter("JSON"),
		zap.WithFile(true, "./logs", "zap_example.log"),
		zap.WithFileLevel("DEBUG"),
		zap.WithFileFormatter("JSON"),
	)

	// Logs básicos
	logger.Info("=== Demonstração Zap ===")
	logger.Info("Logger criado com zap")
	logger.Debug("Debug com zap")
	logger.Warn("Aviso com zap")
	logger.Error("Erro com zap")

	// Campos estruturados
	userLogger := logger.WithFields(
		logr.String("user_id", "12345"),
		logr.String("username", "joao_silva"),
		logr.Bool("active", true),
		logr.Int("login_count", loginCount),
		logr.Float64("score", scoreValue),
		logr.Time("last_login", time.Now()),
		logr.Duration("session_duration", sessionHours),
	)

	userLogger.Info("Usuário logado com sucesso")

	// Campos agrupados
	orderLogger := logger.WithFields(
		logr.Group("order",
			logr.String("id", "ORD-12345"),
			logr.Float64("total", totalAmount),
			logr.String("status", "completed"),
		),
		logr.Group("customer",
			logr.String("name", "Maria Santos"),
			logr.String("email", "maria@exemplo.com"),
		),
	)

	orderLogger.Info("Pedido processado")
}

func demonstrateGlobalLogger() {
	// Configurar logger global
	globalLogger := zap.New(
		zap.WithConsole(true),
		zap.WithConsoleFormatter("JSON"),
	)

	logr.Set(globalLogger)

	// Usar funções globais
	logr.Info("=== Demonstração Logger Global ===")
	logr.Info("Usando logger global")
	logr.Warn("Aviso global")
	logr.Error("Erro global")

	// Com campos
	logr.WithFields(
		logr.String("component", "global"),
		logr.Bool("initialized", true),
	).Info("Sistema inicializado")

	// Logger com campos persistentes
	serviceLogger := logr.WithFields(
		logr.String("service", "api"),
		logr.String("version", "v1.2.3"),
	)

	serviceLogger.Info("Serviço iniciado")
	serviceLogger.WithField(logr.String("endpoint", "/users")).Info("Endpoint registrado")
}

func demonstrateContextUsage() {
	logger := slog.New(
		slog.WithConsole(true),
		slog.WithConsoleFormatter("JSON"),
	)

	logger.Info("=== Demonstração Uso com Contexto ===")

	// Simular uma requisição HTTP
	ctx := context.Background()

	// Adicionar informações da requisição ao contexto
	requestLogger := logger.WithFields(
		logr.String("request_id", "req-12345"),
		logr.String("method", "GET"),
		logr.String("path", "/api/users"),
		logr.String("ip", "192.168.1.100"),
	)

	ctx = requestLogger.ToContext(ctx)

	// Em uma função que recebe o contexto
	processRequest(ctx, "user-123")
}

func processRequest(ctx context.Context, userID string) {
	// Recuperar logger do contexto
	logger := logr.FromContext(ctx)

	// Adicionar informações específicas da função
	userLogger := logger.WithFields(
		logr.String("user_id", userID),
		logr.String("function", "processRequest"),
	)

	userLogger.Info("Processando requisição do usuário")

	// Simular processamento
	time.Sleep(sleepDuration)

	userLogger.WithFields(
		logr.Duration("processing_time", processingDuration),
		logr.Bool("success", true),
	).Info("Requisição processada com sucesso")
}
