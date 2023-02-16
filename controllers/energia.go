package controllers

import (
	"github.com/gin-gonic/gin"
	"imoveis/databases"
	"imoveis/models"
	"imoveis/utils"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

func Energia(c *gin.Context) {

	err := c.Request.ParseForm()
	if err != nil {
		log.Println("Erro ao fazer o ParseForm de energia", err)
	}
	if len(c.PostForm("id")) > 0 {
		alteraEnergiaDoMes(c)
	} else if len(c.PostForm("relogio_1")) > 0 && len(c.PostForm("relogio_2")) > 0 && len(c.PostForm("relogio_3")) > 0 {
		salvaEnergiaDoMes(c)
	}

	dadosTabelaEnergia := montaTabelaEnergia()
	c.HTML(http.StatusOK, "views/energia.html", gin.H{
		"energia": dadosTabelaEnergia,
	})
}

func alteraEnergiaDoMes(c *gin.Context) {
	id, _ := strconv.ParseUint(c.PostForm("id"), 10, 36)
	var energia models.Energia
	databases.DB.First(&energia, id)
	energiaAtualizado := models.Energia{
		ID:         uint(id),
		UpdatedAt:  time.Now(),
		Data:       utils.ParseDate(c.PostForm("data")),
		Relogio1:   utils.StringToInt(c.PostForm("relogio_1")),
		Relogio2:   utils.StringToInt(c.PostForm("relogio_2")),
		Relogio3:   utils.StringToInt(c.PostForm("relogio_3")),
		ValorKwh:   utils.StringToFloat(c.PostForm("valor_kwh")),
		ValorConta: utils.StringToFloat(c.PostForm("valor_conta")),
	}
	databases.DB.Save(&energiaAtualizado)
}

func montaTabelaEnergia() []models.Energia {
	var ultimoRegistro models.Energia
	databases.DB.Order("id desc").Limit(1).Find(&ultimoRegistro)

	var quatroUltimosRegistros []models.Energia
	databases.DB.Order("id desc").Limit(4).Find(&quatroUltimosRegistros)

	var ultimo bool

	sort.Slice(quatroUltimosRegistros, func(i, j int) bool {
		return quatroUltimosRegistros[i].ID < quatroUltimosRegistros[j].ID
	})

	cont := 0
	var dadosTabelaEnergia []models.Energia
	for _, e := range quatroUltimosRegistros {

		if cont > 0 {
			kwhKitnet1 := e.Relogio1 - quatroUltimosRegistros[cont-1].Relogio1
			kwhKitnet2 := e.Relogio2 - quatroUltimosRegistros[cont-1].Relogio2
			kwhKitnet3 := e.Relogio3 - quatroUltimosRegistros[cont-1].Relogio3
			kwhTotal := kwhKitnet1 + kwhKitnet2 + kwhKitnet3
			kitnet1 := float64(kwhKitnet1) / float64(kwhTotal) * e.ValorConta
			kitnet2 := float64(kwhKitnet2) / float64(kwhTotal) * e.ValorConta
			kitnet3 := float64(kwhKitnet3) / float64(kwhTotal) * e.ValorConta

			if e.ID == ultimoRegistro.ID {
				ultimo = true
			}

			dados := models.Energia{
				ID:             e.ID,
				Data:           e.Data,
				Relogio1:       e.Relogio1,
				ValorConta1:    kitnet1,
				Relogio2:       e.Relogio2,
				ValorConta2:    kitnet2,
				Relogio3:       e.Relogio3,
				ValorConta3:    kitnet3,
				ValorKwh:       e.ValorKwh,
				ValorConta:     e.ValorConta,
				UltimoRegistro: ultimo,
			}
			dadosTabelaEnergia = append(dadosTabelaEnergia, dados)
		}
		cont++
	}
	return dadosTabelaEnergia
}

func salvaEnergiaDoMes(c *gin.Context) {
	var energia models.Energia
	databases.DB.Last(&energia)

	data := utils.ParseDate(c.PostForm("data"))
	relogio1 := utils.StringToInt(c.PostForm("relogio_1"))
	relogio2 := utils.StringToInt(c.PostForm("relogio_2"))
	relogio3 := utils.StringToInt(c.PostForm("relogio_3"))
	valorKwh := utils.StringToFloat(c.PostForm("valor_kwh"))
	valorConta := utils.StringToFloat(c.PostForm("valor_conta"))

	if !(relogio1 == energia.Relogio1 && relogio2 == energia.Relogio2 && relogio3 == energia.Relogio3) {
		energiaDoMes := models.Energia{
			Data:       data,
			Relogio1:   relogio1,
			Relogio2:   relogio2,
			Relogio3:   relogio3,
			ValorKwh:   valorKwh,
			ValorConta: valorConta,
		}
		databases.DB.Save(&energiaDoMes)
	} else {
		log.Println("Tentativa de inserir dados repetidos.")
	}
}

func EnergiaForm(c *gin.Context) {
	id := c.Query("editar")
	var energia models.Energia
	databases.DB.First(&energia, id)
	if len(id) > 0 {
		c.HTML(http.StatusOK, "views/energia_form.html", gin.H{
			"energia": energia,
		})
	} else {
		c.HTML(http.StatusOK, "views/energia_form.html", nil)
	}
}
