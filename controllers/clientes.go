package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"imoveis/databases"
	"imoveis/models"
	"imoveis/utils"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func InputSelectClientes(id int) string {
	var clientes []models.Cliente
	databases.DB.Order("nome").Find(&clientes)
	selectClientes := "<select class='form-select' name='cliente' id='cliente' required><option value=''>&nbsp;</option>"
	for _, cliente := range clientes {
		if id == int(cliente.ID) {
			selectClientes +=
				fmt.Sprintf("<option value='%d' selected>%s</option>", cliente.ID, cliente.Nome)
		} else {
			selectClientes +=
				fmt.Sprintf("<option value='%d'>%s</option>", cliente.ID, cliente.Nome)
		}
	}
	selectClientes += "</select>"
	return selectClientes
}

func insereCliente(dadosFormulario url.Values) {
	cliente := models.Cliente{
		Nome:           dadosFormulario["nome"][0],
		DataNascimento: utils.ParseDate(dadosFormulario["data_nascimento"][0]),
		Ci:             dadosFormulario["ci"][0],
		Cpf:            dadosFormulario["cpf"][0],
		Telefone1:      dadosFormulario["telefone_1"][0],
		Telefone2:      dadosFormulario["telefone_2"][0],
	}
	databases.DB.Save(&cliente)
}

func atualizaCliente(dadosFormulario url.Values, id string) {
	var cliente models.Cliente
	databases.DB.Find(&cliente, id)
	cliente.Nome = dadosFormulario["nome"][0]
	cliente.DataNascimento = utils.ParseDate(dadosFormulario["data_nascimento"][0])
	cliente.Ci = dadosFormulario["ci"][0]
	cliente.Cpf = dadosFormulario["cpf"][0]
	cliente.Telefone1 = dadosFormulario["telefone_1"][0]
	cliente.Telefone2 = dadosFormulario["telefone_2"][0]
	databases.DB.Save(&cliente)
}

func apagaCliente(c *gin.Context) {
	var cliente models.Cliente
	id := func() int {
		id, err := strconv.Atoi(c.Query("apagar"))
		if err != nil {
			log.Panic("Não conseguimos converter o parâmetro de 'apagar'.")
		}
		return id
	}()
	databases.DB.First(&cliente, id)
	databases.DB.Delete(&cliente, id)
}

func montaTabelaDeClientes() []models.Cliente {
	var clientes []models.Cliente
	databases.DB.Order("nome").Find(&clientes)

	var imoveis []models.Imovel
	var clientesParaMontarTabela []models.Cliente
	var dadosDeUmCliente models.Cliente
	for _, c := range clientes {
		var ids string
		databases.DB.Where("cliente_id = ?", c.ID).Find(&imoveis)
		for _, imovel := range imoveis {
			if int(c.ID) == imovel.ClienteID {
				ids += fmt.Sprintf("| %s %s |", imovel.Tipo, imovel.Numero)
			}
		}
		dadosDeUmCliente = models.Cliente{
			ID:             c.ID,
			Nome:           c.Nome,
			Locacoes:       ids,
			DataNascimento: c.DataNascimento,
			Ci:             c.Ci,
			Cpf:            c.Cpf,
			Telefone1:      c.Telefone1,
			Telefone2:      c.Telefone2,
		}
		clientesParaMontarTabela = append(clientesParaMontarTabela, dadosDeUmCliente)
	}
	return clientesParaMontarTabela
}

func Clientes(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		log.Println("erro ao fazer o ParseForm de clientes", err)
	}
	id := c.Request.PostForm.Get("id")
	nome := c.Request.PostForm.Get("nome")
	apagar := c.Query("apagar")

	demandaInserirCliente := len(id) == 0 && len(nome) != 0
	demandaAtualizarCliente := len(id) > 0
	demandaApagarCliente := len(apagar) > 0

	if demandaInserirCliente {
		insereCliente(c.Request.PostForm)
	} else if demandaAtualizarCliente {
		atualizaCliente(c.Request.PostForm, id)
	} else if demandaApagarCliente {
		apagaCliente(c)
	}

	tabelaDeClientes := montaTabelaDeClientes()
	c.HTML(http.StatusOK, "views/clientes.html", gin.H{
		"clientes": tabelaDeClientes,
	})
}

func ClientesForm(c *gin.Context) {
	id := c.Query("editar")
	if len(id) > 0 {
		var cliente models.Cliente
		databases.DB.Find(&cliente, id)
		c.HTML(http.StatusOK, "views/clientes_form.html", gin.H{
			"cliente": cliente,
		})
	} else {
		c.HTML(http.StatusOK, "views/clientes_form.html", nil)
	}
}
