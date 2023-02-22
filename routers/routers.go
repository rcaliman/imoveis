package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"imoveis/auth"
	"imoveis/controllers"
	"imoveis/utils"
	"os"
)

func IniciaRoteamento() {
	port := os.Getenv("PORT")
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"extraiData":       utils.ExtraiData,
		"formataData":      utils.FormataData,
		"formataDinheiro":  utils.FormataDinheiro,
		"mascaraDocumento": utils.MascaraDocumento,
		"floatToString":    utils.FloatToString,
		"arredonda2":       utils.Arredonda2,
	})
	r.LoadHTMLGlob("templates/*/**")
	r.Static("/css", "./static/css")
	r.Static("/img", "./static/img")
	r.Static("/js", "./static/js")

	r.GET("/", controllers.Home)
	sistema := r.Group("/sistema", gin.BasicAuth(auth.AuthUsuarios()))
	sistema.POST("/imoveis", controllers.Imoveis)
	sistema.GET("/imoveis", controllers.Imoveis)
	sistema.GET("/imoveis/form", controllers.ImoveisForm)
	sistema.POST("/clientes", controllers.Clientes)
	sistema.GET("/clientes", controllers.Clientes)
	sistema.GET("/clientes/form", controllers.ClientesForm)
	sistema.POST("/energia", controllers.Energia)
	sistema.GET("/energia", controllers.Energia)
	sistema.GET("/energia/form", controllers.EnergiaForm)
	sistema.POST("/recibos", controllers.Recibos)
	sistema.GET("/usuarios", controllers.Usuarios)
	sistema.GET("/usuarios/form", controllers.UsuariosForm)
	sistema.POST("/usuarios/form", controllers.UsuariosForm)
	sistema.POST("/usuarios", controllers.Usuarios)

	r.Run(fmt.Sprintf(":%s", port))
}
