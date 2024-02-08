package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"snippetbox.edgardev.net/internal/models"
)

// Para utilizar el logger que has creado en todas tus dependencias, haras uso
// de Dependency injection

type Application struct {
	logger *slog.Logger
	snippets *models.SnippetModel

}

func main() {

	
	// Creando un command line flag, sirve para enviar variables desde la terminal
	addr := flag.String("addr", ":4000", "HTTP NETWORK ADDRESS")
	dsn := flag.String("dsn", "web:root@/snippetbox?parseTime=True", "This is the database path string")

	// crendo un log personalizado
	loggerHandler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(loggerHandler)

	db, err := databaseConnection(*dsn)
	
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// cerramos la base de daots
	defer db.Close()
	
	// creas una instancia para decirle a la struct con que logger se tiene que vincular
	app := &Application{
		logger: logger,
		//         instanciando la clase SnippetModel
		snippets: &models.SnippetModel{DB: db},
	}

	logger.Info("Server is runnin on port", "addr", *addr)


	// Antes de usar las comman line flags debes convertirlas
	flag.Parse()

	// starting the server using the http package
	// al usar una command line flag recuerda que debes pasar el *puntero, no la variable como tal
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1) // if an error occurs the server stop running
}


// @param dsn -> this is the database path
func databaseConnection(dsn string)(*sql.DB, error){

	db, err:= sql.Open("mysql",dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	
	if err != nil {
	db.Close()
	return nil, err
	}
	
	return db, nil
}

