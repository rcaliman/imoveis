package sistemaVelho

import (
	novoDB "imoveis/databases"
	velhoDB "imoveis/databases/sistemaVelho"
	novoMD "imoveis/models"
	velhoMD "imoveis/models/sistemaVelho"
	"strings"
	"time"
)

func parseDate(t time.Time) time.Time {
	ano := t.Year()
	dia := t.Day()
	data := time.Date(ano, t.Month(), dia, 0, 0, 0, 0, time.Local)
	return data
}

func stripDocs(doc string) string {
	doc = strings.ReplaceAll(doc, "/", "")
	doc = strings.ReplaceAll(doc, "-", "")
	doc = strings.ReplaceAll(doc, ".", "")
	return doc
}

func MigraAlugueis() {
	velhoDB.ConectaDB()

	var sistemaNovoClientes []novoMD.Cliente
	novoDB.DB.Find(&sistemaNovoClientes)

	var sistemaVelhoClientes []velhoMD.Al_Cliente
	velhoDB.DB.Find(&sistemaVelhoClientes)

	if len(sistemaNovoClientes) == 0 && len(sistemaVelhoClientes) > 0 {

		var sistemaVelhoClientes []velhoMD.Al_Cliente
		velhoDB.DB.Find(&sistemaVelhoClientes)

		var sistemaNovoCliente novoMD.Cliente
		for _, cliente := range sistemaVelhoClientes {
			sistemaNovoCliente = novoMD.Cliente{
				Nome:           cliente.Nome,
				DataNascimento: parseDate(cliente.DataNasc),
				Ci:             cliente.Ci,
				Cpf:            stripDocs(cliente.Cpf),
				Telefone1:      cliente.Tel1,
				Telefone2:      cliente.Tel2,
			}
			novoDB.DB.Create(&sistemaNovoCliente)

		}

		var sistemaVelhoImoveis []velhoMD.Al_Imoveis
		velhoDB.DB.Find(&sistemaVelhoImoveis)

		for _, imovel := range sistemaVelhoImoveis {
			var cliente novoMD.Cliente
			novoDB.DB.Where("nome like ?", imovel.Cliente).First(&cliente)
			sistemaNovoImovel := novoMD.Imovel{
				Tipo:         strings.ToLower(imovel.Tipo),
				Numero:       imovel.Numero,
				Local:        imovel.Local,
				ClienteID:    int(cliente.ID),
				ValorAluguel: imovel.ValorAluguel,
				Observacao:   imovel.Observacao,
				DiaBase:      imovel.DiaBase,
			}
			novoDB.DB.Create(&sistemaNovoImovel)
		}

		var sistemaVelhoEnergia []velhoMD.Al_Energia
		velhoDB.DB.Order("data").Find(&sistemaVelhoEnergia)

		for _, energia := range sistemaVelhoEnergia {
			sistemaNovoEnergia := novoMD.Energia{
				Data:       parseDate(energia.Data),
				Relogio1:   energia.Relogio1,
				Relogio2:   energia.Relogio2,
				Relogio3:   energia.Relogio3,
				ValorKwh:   energia.ValorKwh,
				ValorConta: energia.ValorConta,
			}
			novoDB.DB.Create(&sistemaNovoEnergia)
		}
	}
}
