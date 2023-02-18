package sistemaVelho

import "time"

type Al_Cliente struct {
	Nome     string    `json:"nome""`
	DataNasc time.Time `json:"data_nasc"`
	Ci       string    `json:"string"`
	Cpf      string    `json:"string"`
	Tel1     string    `json:"string"`
	Tel2     string    `json:"string"`
}
