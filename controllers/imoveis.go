package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"imoveis/databases"
	"imoveis/models"
	"imoveis/utils"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

func inputSelectTipo(id string) string {
	var imovel models.Imovel
	databases.DB.First(&imovel, id)
	tipos := []string{"apartamento", "condominio", "kitnet", "loja", "sala comercial"}
	selectTipo := fmt.Sprintf("<select id='tipo' name='tipo' class='form-select'>")
	for _, tipo := range tipos {
		if tipo == imovel.Tipo {
			selectTipo += fmt.Sprintf("<option selected>%s</option>", tipo)
		} else {
			selectTipo += fmt.Sprintf("<option>%s</option>", tipo)
		}
	}
	selectTipo += "</select>"
	return selectTipo
}

func insereImovel(dadosFormulario url.Values) {
	imovel := models.Imovel{
		Tipo:         dadosFormulario["tipo"][0],
		Numero:       dadosFormulario["numero"][0],
		Local:        dadosFormulario["local"][0],
		ClienteID:    utils.StringToInt(dadosFormulario["cliente"][0]),
		ValorAluguel: utils.StringToFloat(dadosFormulario["valor_aluguel"][0]),
		Observacao:   dadosFormulario["observacao"][0],
		DiaBase:      utils.StringToInt(dadosFormulario["dia_base"][0]),
	}
	databases.DB.Create(&imovel)
}

func atualizaImovel(dadosFormulario url.Values, id string) {
	var imovel models.Imovel
	databases.DB.First(&imovel, id)
	imovel.Tipo = dadosFormulario["tipo"][0]
	imovel.Numero = dadosFormulario["numero"][0]
	imovel.Local = dadosFormulario["local"][0]
	imovel.ClienteID = utils.StringToInt(dadosFormulario["cliente"][0])
	imovel.ValorAluguel = utils.StringToFloat(dadosFormulario["valor_aluguel"][0])
	imovel.Observacao = dadosFormulario["observacao"][0]
	imovel.DiaBase = utils.StringToInt(dadosFormulario["dia_base"][0])
	databases.DB.Save(&imovel)
}

func apagaImovel(id string) {
	var imovel models.Imovel
	databases.DB.Delete(&imovel, id)
}

func ordenaTabelaImoveis(imoveis []models.Imovel, ordenador string) {
	switch ordenador {
	case "tipo":
		sort.Slice(imoveis, func(i, j int) bool {
			comparativo := strings.Compare(imoveis[j].Tipo, imoveis[i].Tipo)
			if comparativo > 0 {
				return true
			} else {
				return false
			}
		})
	case "numero":
		sort.Slice(imoveis, func(i, j int) bool {
			numero1, _ := strconv.Atoi(imoveis[j].Numero)
			numero2, _ := strconv.Atoi(imoveis[i].Numero)
			return numero1 > numero2
		})
	case "local":
		sort.Slice(imoveis, func(i, j int) bool {
			comparativo := strings.Compare(imoveis[j].Local, imoveis[i].Local)
			if comparativo > 0 {
				return true
			} else {
				return false
			}
		})
	case "valor_aluguel":
		sort.Slice(imoveis, func(i, j int) bool {
			return imoveis[i].ValorAluguel < imoveis[j].ValorAluguel
		})
	case "dia_base":
		sort.Slice(imoveis, func(i, j int) bool {
			return imoveis[i].DiaBase < imoveis[j].DiaBase
		})
	default:
		sort.Slice(imoveis, func(i, j int) bool {
			comparativo := strings.Compare(imoveis[j].Cliente.Nome, imoveis[i].Cliente.Nome)
			if comparativo > 0 {
				return true
			} else {
				return false
			}
		})
	}
}

func Imoveis(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		log.Println("Erro ao fazer o ParseForm de imoveis:", err)
	}
	id := c.PostForm("id")
	numero := c.PostForm("numero")
	apagar := c.Query("apagar")
	ordenador := c.Query("ordenador")

	demandaInserirImovel := len(id) == 0 && len(numero) > 0
	demandaAtualizarImovel := len(id) > 0
	demandaApagarImovel := len(apagar) > 0

	if demandaInserirImovel {
		insereImovel(c.Request.PostForm)
	} else if demandaAtualizarImovel {
		atualizaImovel(c.Request.PostForm, id)
	} else if demandaApagarImovel {
		apagaImovel(apagar)
	}

	var imoveis []models.Imovel
	databases.DB.Preload("Cliente").Find(&imoveis)

	ordenaTabelaImoveis(imoveis, ordenador)

	c.HTML(http.StatusOK, "views/imoveis.html", gin.H{
		"imoveis":      imoveis,
		"select_meses": template.HTML(utils.InputSelectMeses()),
		"select_anos":  template.HTML(utils.InputSelectAnos()),
	})
}

func ImoveisForm(c *gin.Context) {
	id := c.Query("editar")
	if len(id) > 0 {
		var imovel models.Imovel
		databases.DB.First(&imovel, id)
		c.HTML(http.StatusOK, "views/imoveis_form.html", gin.H{
			"imovel":          imovel,
			"select_clientes": template.HTML(InputSelectClientes(imovel.ClienteID)),
			"select_tipos":    template.HTML(inputSelectTipo(id)),
		})
	} else {
		c.HTML(http.StatusOK, "views/imoveis_form.html", gin.H{
			"select_clientes": template.HTML(InputSelectClientes(0)),
			"select_tipos":    template.HTML(inputSelectTipo("0")),
		})
	}
}
