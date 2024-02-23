package main

import "snippetbox.edgardev.net/internal/models"

type TemplateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
