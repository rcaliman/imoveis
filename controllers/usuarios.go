package controllers

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"imoveis/databases"
	"imoveis/models"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Usuarios(c *gin.Context) {
	err := c.Request.ParseForm()
	user, _, _ := c.Request.BasicAuth()
	if err != nil {
		log.Println("Problema ao processar formulário: ", err)
	}
	if c.PostForm("usuario") == user {
		alteraSenha(c)
	}

	var usuario models.Usuario
	databases.DB.Where("usuario = ?", user).First(&usuario)
	tipo := make(map[string]bool)
	if usuario.Tipo == "administrador" {
		tipo["administrador"] = true
	} else {
		tipo["administrador"] = false
	}
	c.HTML(http.StatusOK, "views/usuarios.html", gin.H{
		"usuario": usuario,
		"tipo":    tipo,
	})
}

func InputSelectTipoUsuario(tipo string) string {
	tipoUsuario := []string{"administrador", "operador"}
	selectTipoUsuario := fmt.Sprintf("<select name='tipo' id='tipo' class='form-select'>")
	for _, t := range tipoUsuario {
		if tipo != t {
			selectTipoUsuario += fmt.Sprintf("<option>%s</option>", t)
		} else {
			selectTipoUsuario += fmt.Sprintf("<option selected>%s</option>", t)
		}
	}
	selectTipoUsuario += fmt.Sprintf("</select>")
	return selectTipoUsuario
}

func alteraSenha(c *gin.Context) {
	mapUsuario := map[string]string{
		"usuario":           strings.TrimSpace(c.PostForm("usuario")),
		"senhaatual":        strings.TrimSpace(c.PostForm("senhaatual")),
		"novasenha":         strings.TrimSpace(c.PostForm("novasenha")),
		"confirmanovasenha": strings.TrimSpace(c.PostForm("confirmanovasenha")),
	}
	var usuarioSalvo models.Usuario

	databases.DB.Where("usuario = ?", mapUsuario["usuario"]).First(&usuarioSalvo)

	formSenhaAtual := base64.StdEncoding.EncodeToString([]byte(mapUsuario["senhaatual"]))
	formNovaSenha := base64.StdEncoding.EncodeToString([]byte(mapUsuario["novasenha"]))

	if usuarioSalvo.Senha == formSenhaAtual {
		if strings.Compare(mapUsuario["novasenha"], mapUsuario["confirmanovasenha"]) == 0 {
			usuarioSalvo.Senha = formNovaSenha
			databases.DB.Save(&usuarioSalvo)
		}
	}
}

func UsuariosForm(c *gin.Context) {
	if err := c.Request.ParseForm(); err == nil {
		id := c.PostForm("id")
		processaForm(c, id)

	} else {
		log.Println("Erro ao processar o formulário", err)
	}

	editar := c.Query("editar")
	apagar := c.Query("apagar")
	user, _, _ := c.Request.BasicAuth()

	var usuario models.Usuario
	databases.DB.Where("usuario = ?", user).First(&usuario)

	if usuario.Tipo == "administrador" {
		if len(editar) > 0 {
			var usuarioEditar models.Usuario
			databases.DB.First(&usuarioEditar, editar)
			var usuarios []models.Usuario
			databases.DB.Find(&usuarios)
			c.HTML(http.StatusOK, "views/usuarios_form.html", gin.H{
				"usuarios":            usuarios,
				"usuario":             usuarioEditar,
				"select_tipo_usuario": template.HTML(InputSelectTipoUsuario(usuarioEditar.Tipo)),
			})
		} else if len(apagar) > 0 {
			var usuarioApagar models.Usuario
			databases.DB.Unscoped().Delete(&usuarioApagar, apagar)

			var usuarios []models.Usuario
			databases.DB.Find(&usuarios)

			c.HTML(http.StatusOK, "views/usuarios_form.html", gin.H{
				"usuarios":            usuarios,
				"select_tipo_usuario": template.HTML(InputSelectTipoUsuario("")),
			})
		} else {
			var usuarios []models.Usuario
			databases.DB.Find(&usuarios)

			c.HTML(http.StatusOK, "views/usuarios_form.html", gin.H{
				"usuarios":            usuarios,
				"select_tipo_usuario": template.HTML(InputSelectTipoUsuario("")),
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "você não é administrador",
		})
	}
}

func processaForm(c *gin.Context, id string) {
	usuario := strings.TrimSpace(c.PostForm("usuario"))
	tipo := strings.TrimSpace(c.PostForm("tipo"))
	senha := strings.TrimSpace(c.PostForm("senha"))
	confirmasenha := strings.TrimSpace(c.PostForm("confirmasenha"))
	if len(usuario) > 0 && len(tipo) > 0 && len(senha) > 0 && len(confirmasenha) > 0 {
		if senha == confirmasenha {
			usuario := models.Usuario{
				Usuario: usuario,
				Tipo:    tipo,
				Senha:   base64.StdEncoding.EncodeToString([]byte(senha)),
			}
			if len(id) == 0 {
				databases.DB.Create(&usuario)
			} else if len(id) > 0 {
				usuario.ID = func() uint {
					i, _ := strconv.Atoi(id)
					return uint(i)
				}()
				databases.DB.Save(&usuario)
			}
		} else {
			log.Println("As senhas são diferentes")
		}
	}
}
