package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}


type SnippetModel struct {
	DB * sql.DB
}


func (model SnippetModel)Insert(title string, content string, expires int)(int, error){
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

func (model SnippetModel)GetById(id int)(Snippet, error){
	return Snippet{}, nil
}

// This will return the 10 most recently snippets
func (model SnippetModel) Lastest()([]Snippet, error){
	return []Snippet{}, nil
}