package models

import (
	"errors"
)

// La razon por la cual creamos este error personalizado en lugar de retornarl el error SQL del modelo
// es por que asi encapsulamos por completo todo lo que tenga que ver con la base de datos
// dentro del modelo y no sale del modelo nada del SQL
var ErrNoRecord = errors.New("models: no matching record found :c .sad")
