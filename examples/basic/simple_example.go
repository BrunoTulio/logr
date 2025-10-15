package main

import (
	"time"

	"github.com/BrunoTulio/logr"
	"github.com/BrunoTulio/logr/adapters/slog.v1"
	"github.com/BrunoTulio/logr/adapters/zap.v1"
)

func main() {
	// Exemplo 1: Logger básico com slog
	basicSlogExample()

	// Exemplo 2: Logger básico com zap
	basicZapExample()

	// Exemplo 3: Logger com campos estruturados
	structuredFieldsExample()

	// Exemplo 4: Logger com arquivo
	fileLoggerExample()
}

func basicSlogExample() {
	logger := slog.New(
		slog.WithConsole(true),
		slog.WithConsoleLevel("INFO"),
		slog.WithConsoleFormatter("TEXT"),
	)

	logger.Info("=== Exemplo Básico com Slog ===")
	logger.Info("Aplicação iniciada")
	logger.Warn("Aviso importante")
	logger.Error("Erro encontrado")
}

func basicZapExample() {
	logger := zap.New(
		zap.WithConsole(true),
		zap.WithConsoleLevel("INFO"),
		zap.WithConsoleFormatter("JSON"),
	)

	logger.Info("=== Exemplo Básico com Zap ===")
	logger.Info("Sistema iniciado")
	logger.Warn("Aviso crítico")
	logger.Error("Falha no sistema")
}

func structuredFieldsExample() {
	logger := slog.New(
		slog.WithConsole(true),
		slog.WithConsoleFormatter("JSON"),
	)

	logger.Info("=== Exemplo com Campos Estruturados ===")

	// Campos simples
	logger.WithFields(
		logr.String("user_id", "12345"),
		logr.Bool("active", true),
		logr.Int("age", 30),
	).Info("Usuário logado")

	// Campos agrupados
	logger.WithFields(
		logr.Group("user",
			logr.String("name", "João Silva"),
			logr.String("email", "joao@exemplo.com"),
		),
		logr.Group("order",
			logr.String("id", "ORD-12345"),
			logr.Float64("total", 299.99),
		),
	).Info("Pedido processado")

	// Campos com tipos diferentes
	logger.WithFields(
		logr.Uint64("id", 123456789),
		logr.Float64("score", 95.7),
		logr.Time("created_at", time.Now()),
		logr.Duration("duration", time.Hour*2),
	).Info("Dados processados")
}

func fileLoggerExample() {
	logger := zap.New(
		zap.WithConsole(true),
		zap.WithConsoleFormatter("JSON"),
		zap.WithFile(true, "./logs", "example.log"),
		zap.WithFileFormatter("JSON"),
		zap.WithFileRotation(10, 7, true), // 10MB, 7 dias, comprimir
	)

	logger.Info("=== Exemplo com Arquivo ===")
	logger.Info("Log salvo em arquivo e console")
	
	logger.WithFields(
		logr.String("service", "example"),
		logr.String("version", "1.0.0"),
	).Info("Serviço configurado")
}
