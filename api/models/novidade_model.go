package models

type NovidadeModel struct {
	Titulo string `json:"titulo"`
	Notas  string `json:"notas"`
	Versao string `json:"versao"`
}

var Novidades []NovidadeModel
