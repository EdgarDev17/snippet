package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// Para utilizar el logger que has creado en todas tus dependencias, haras uso
// de Dependency injection

type Application struct {
	logger *slog.Logger
}

func main() {

	// crendo un log personalizado
	loggerHandler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(loggerHandler)

	// creas una instancia para decirle a la struct con que logger se tiene que vincular
	app := &Application{
		logger: logger,
	}

	// Creando un command line flag, sirve para enviar variables desde la terminal
	// cuando se inicia una aplicacion Go
	addr := flag.String("addr", ":4000", "HTTP NETWORK ADDRESS")

	logger.Info("Server is runnin on", "addr", *addr)

	// Antes de usar las comman line flags debes convertirlas
	flag.Parse()

	// starting the server using the http package
	// al usar una command line flag recuerda que debes pasar el *puntero, no la variable como tal
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1) // if an error occurs the server stop running
}
