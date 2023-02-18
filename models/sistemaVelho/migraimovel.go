package sistemaVelho

type Al_Imoveis struct {
	Tipo         string  `json:"tipo"`
	Numero       string  `json:"numero"`
	Local        string  `json:"local"`
	Cliente      string  `json:"cliente"`
	ValorAluguel float64 `json:"valor_aluguel" gorm:"type:double"`
	Observacao   string  `json:"observacao"`
	DiaBase      int     `json:"dia_base"`
}
