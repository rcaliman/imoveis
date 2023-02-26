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
	"sort"
	"strings"
)

func Recibos(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		log.Println("Erro ao fazer o ParseForm de recibos", err)
	}
	ids := c.Request.Form["imprimir"]
	var imoveis []models.Imovel
	databases.DB.Preload("Cliente").Preload("Imovel").Where("id in ?", ids).Find(&imoveis)
	sort.Slice(imoveis, func(i, j int) bool {
		if strings.Compare(imoveis[i].Cliente.Nome, imoveis[j].Cliente.Nome) < 0 {
			return true
		} else {
			return false
		}
	})
	var recibo string
	cont := 1
	for _, imovel := range imoveis {
		if imovel.Tipo == "condominio" {
			recibo = reciboCondominio(c, recibo, imovel, cont)
		} else {
			recibo = reciboAluguel(c, recibo, imovel, cont)
		}
		cont++
	}
	c.HTML(http.StatusOK, "views/recibos.html", gin.H{
		"recibo": template.HTML(recibo),
	})
}

func reciboCondominio(c *gin.Context, recibo string, imovel models.Imovel, cont int) string {
	qt := 0
	for qt < 2 {
		mesPassado, anoPassado := utils.DizMesAnterior(c.PostForm("recibo_mes"), c.PostForm("recibo_ano"))
		recibo += "<div class='recibo'>"
		recibo += "<h1 class='titulo'>RECIBO DE CONDOMÍNIO</h1>"
		recibo += fmt.Sprintf(
			"<div id='linharecibo' class='linharecibo'>Recebi de <b>%s</b> a importância de <b><span id='valorExtenso%d-%d'></span></b> referente ao condominio do mês de <b>%s</b> de <b>%s</b> de %s no Edifício Caliman.</div>", imovel.Cliente.Nome, imovel.ID, qt, c.PostForm("recibo_mes"), c.PostForm("recibo_ano"), imovel.Observacao)
		recibo += fmt.Sprintf(
			"<p class='linhadata'>Colatina-ES, 1 de %v de %v.", mesPassado, anoPassado)
		recibo += fmt.Sprintf(
			"<p class='linhaassinatura'>___________________________________<br>Darci Francisco Caliman<br>Proprietário</p>")
		recibo += fmt.Sprintf(
			"<p class='linhatelefone'>&nbsp;%s %s</p>", imovel.Cliente.Telefone1, imovel.Cliente.Telefone2)
		recibo += fmt.Sprintf(
			"</div><hr style='border-top: solid 2px;'>")
		recibo += fmt.Sprintf(
			"<script>document.getElementById('valorExtenso%d-%d').innerHTML = extenso('%g', {mode: 'currency'});</script>", imovel.ID, qt, imovel.ValorAluguel)
		qt += 1
	}
	recibo += fmt.Sprintf("<p class='contapagina' style='page-break-after: always'>%d</p>", cont)
	return recibo
}

func reciboAluguel(c *gin.Context, recibo string, imovel models.Imovel, cont int) string {
	qt := 0
	for qt < 2 {
		recibo += "<div class='recibo'>"
		recibo += "<h1 class='titulo'>RECIBO</h1>"
		recibo += fmt.Sprintf(
			"<div id='linharecibo' class='linharecibo'>Recebi de <b>%s</b> a importância de <b><span id='valorExtenso%d-%d'></span></b> referente ao aluguel do(a) <b>%s</b> numero <b>%s</b>. * * * * * * * *</div>", imovel.Cliente.Nome, imovel.ID, qt, imovel.Tipo, imovel.Numero)
		recibo += fmt.Sprintf(
			"<p class='linhadata'>Colatina-ES, %s de %s de %s.", utils.IntToString(imovel.DiaBase), c.PostForm("recibo_mes"), c.PostForm("recibo_ano"))
		recibo += fmt.Sprintf(
			"<p class='linhaassinatura'>___________________________________<br>Darci Francisco Caliman<br>Proprietário</p>")
		recibo += fmt.Sprintf(
			"<p class='linhatelefone'>&nbsp;%s %s</p>", imovel.Cliente.Telefone1, imovel.Cliente.Telefone2)
		recibo += fmt.Sprintf(
			"</div><hr style='border-top: solid 2px;'>")
		recibo += fmt.Sprintf(
			"<script>document.getElementById('valorExtenso%d-%d').innerHTML = extenso('%g', {mode: 'currency'});</script>", imovel.ID, qt, imovel.ValorAluguel)
		qt += 1
	}
	recibo += fmt.Sprintf("<p class='contapagina' style='page-break-after: always'>%d</p>", cont)
	return recibo
}
