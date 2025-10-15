package main

import (
	"github.com/BrunoTulio/logr"
	"github.com/BrunoTulio/logr/adapters/slog.v1"
	"github.com/BrunoTulio/logr/adapters/zap.v1"
)

// Exemplos de configuração para diferentes ambientes

// Configuração para Desenvolvimento
func developmentConfig() logr.Logger {
	return slog.New(
		slog.WithConsole(true),
		slog.WithConsoleLevel("DEBUG"),
		slog.WithConsoleFormatter("TEXT"), // Mais legível para desenvolvimento
		slog.WithAddSource(true),          // Mostrar arquivo e linha
	)
}

// Configuração para Produção
func productionConfig() logr.Logger {
	return zap.New(
		zap.WithConsole(true),
		zap.WithConsoleLevel("INFO"),
		zap.WithConsoleFormatter("JSON"), // Estruturado para produção

		zap.WithFile(true, "/var/log/app", "application.log"),
		zap.WithFileLevel("DEBUG"), // Logs mais detalhados em arquivo
		zap.WithFileFormatter("JSON"),
		zap.WithFileRotation(100, 30, true), // 100MB, 30 dias, comprimir
	)
}

// Configuração para Testes
func testConfig() logr.Logger {
	return slog.New(
		slog.WithConsole(false), // Sem output no console durante testes
		// Sem arquivo também - logs descartados
	)
}

// Configuração para Docker/Containers
func dockerConfig() logr.Logger {
	return zap.New(
		zap.WithConsole(true),
		zap.WithConsoleLevel("INFO"),
		zap.WithConsoleFormatter("JSON"), // Para coleta de logs estruturados

		// Sem arquivo - logs vão para stdout/stderr do container
	)
}

// Configuração para Microserviços
func microserviceConfig(serviceName string) logr.Logger {
	return zap.New(
		zap.WithConsole(true),
		zap.WithConsoleLevel("INFO"),
		zap.WithConsoleFormatter("JSON"),

		zap.WithFile(true, "/var/log/"+serviceName, serviceName+".log"),
		zap.WithFileLevel("DEBUG"),
		zap.WithFileFormatter("JSON"),
		zap.WithFileRotation(50, 7, true), // 50MB, 7 dias para microserviços
	).WithFields(
		logr.String("service", serviceName),
		logr.String("environment", "production"),
	)
}

// Configuração para Alta Performance
func highPerformanceConfig() logr.Logger {
	return zap.New(
		zap.WithConsole(true),
		zap.WithConsoleLevel("WARN"), // Apenas warnings e erros no console
		zap.WithConsoleFormatter("JSON"),

		zap.WithFile(true, "/var/log/app", "app.log"),
		zap.WithFileLevel("INFO"),
		zap.WithFileFormatter("JSON"),
		zap.WithFileRotation(200, 14, true), // Arquivos maiores, menos rotação
	)
}

// Configuração para Debugging
func debugConfig() logr.Logger {
	return slog.New(
		slog.WithConsole(true),
		slog.WithConsoleLevel("DEBUG"),
		slog.WithConsoleFormatter("TEXT"),
		slog.WithAddSource(true),

		slog.WithFile(true, "./debug", "debug.log"),
		slog.WithFileLevel("DEBUG"),
		slog.WithFileFormatter("TEXT"),      // Mais legível para debug
		slog.WithFileRotation(10, 1, false), // Arquivos pequenos, sem compressão
	)
}

func main() {
	// Exemplos de uso das configurações

	// Desenvolvimento
	devLogger := developmentConfig()
	devLogger.Info("Logger de desenvolvimento configurado")

	// Produção
	prodLogger := productionConfig()
	prodLogger.WithFields(
		logr.String("environment", "production"),
		logr.String("version", "1.0.0"),
	).Info("Aplicação iniciada em produção")

	// Microserviço
	userServiceLogger := microserviceConfig("user-service")
	userServiceLogger.Info("Serviço de usuários iniciado")

	// Debug
	debugLogger := debugConfig()
	debugLogger.Debug("Modo debug ativado")
}
