package models

import (
	"gorm.io/gorm"
)

type Imovel struct {
	gorm.Model
	Tipo         string  `json:"tipo"`
	Numero       string  `json:"numero"`
	Local        string  `json:"local"`
	ClienteID    int     `json:"cliente_id"`
	Cliente      Cliente `json:"cliente"`
	ValorAluguel float64 `json:"valor_aluguel"`
	Observacao   string  `json:"observacao"`
	DiaBase      int     `json:"dia_base"`
}
