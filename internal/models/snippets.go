package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (model SnippetModel) Insert(title string, content string, expires int) (int, error) {
	var sqlQuery string = `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := model.DB.Exec(sqlQuery, title, content, expires)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (model SnippetModel) GetById(id int) (Snippet, error) {

	// creamos la consulta SQL
	sqlQuery := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// Ejecutamos la consulta y retornamos la fila agregada
	row := model.DB.QueryRow(sqlQuery, id)

	// instanciamos la estructura en donde almacenaremos los datos que vienen de la row
	var snippet Snippet

	// usamos row scan para mandar cada columna de nuestra fila (atributos de la tabla) a nuestra instancia
	// de la estructura Snippet, recuerda que usamos punteros para ir directamente a la instancia ya creada
	// en memoria
	err := row.Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)

	// verificamos si existe un error
	if err != nil {
		// errors.Is() compara si los dos errores son los mismos, si err == sql.ErrNoRows
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	return snippet, nil
}

// This will return the 10 most recently snippets
func (model SnippetModel) Lastest() ([]Snippet, error) {

	sqlQuery := `SELECT id, title, content, created, expires FROM snippets
				WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := model.DB.Query(sqlQuery)

	if err != nil {
		return nil, err
	}

	// 🔴 ESTO ES SUMAMENTE IMPORTANTE PARA NO PETAR LA CONEXION A LA BD
	// 🔴 la palabra "defer" lo que hace es esperar hasta que la función termine su proceso
	// 🔴 (justo antes de hacer return) para ejecutar el codigo que le sigue a "defer"
	defer rows.Close()

	// iniicializamos la variable donde iran las rows

	var snippets []Snippet

	//vamos a iterar cada una de las rows
	for rows.Next() {
		var currentSnippet Snippet

		err := rows.Scan(&currentSnippet.ID, &currentSnippet.Title, &currentSnippet.Content,
			&currentSnippet.Created, &currentSnippet.Expires)

		if err != nil {
			return nil, err
		}

		snippets = append(snippets, currentSnippet)
	}

	// cuando haya terminado de iterar cada row de la base de datos, verificamos si existe alguin error
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// If everything went OK then return the Snippets slice.
	return snippets, nil
}
