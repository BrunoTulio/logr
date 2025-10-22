package main

import (
	"github.com/BrunoTulio/logr"
	logrusadapter "github.com/BrunoTulio/logr/adapters/logrus.v1"
)

func main() {
	// Inicializar logger com logrus
	_ = logrusadapter.New(
		logrusadapter.WithConsole(true),
		logrusadapter.WithConsoleLevel("DEBUG"),
		logrusadapter.WithConsoleFormatter("TEXT"),
		logrusadapter.WithFile(true, "./logs", "app.log"),
		logrusadapter.WithFileLevel("INFO"),
		logrusadapter.WithFileFormatter("JSON"),
		logrusadapter.WithFileRotation(10, 7, true),
	)

	// Usar logger global
	logr.Info("Aplicação iniciada")
	logr.Debug("Modo debug ativado")

	// Logger com campos
	logr.WithFields(
		logr.String("user", "john_doe"),
		logr.Int("age", 30),
	).Info("Usuário logado")

	// Logger com campo único
	logr.WithField(logr.String("service", "auth")).Warn("Tentativa de login falhada")

	// Exemplo de erro
	logr.Error("Erro interno do servidor")

	// Logger fatal (termina o programa)
	// logr.Fatal("Erro fatal")
}
