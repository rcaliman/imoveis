package auth

import (
	"encoding/base64"
	"imoveis/databases"
	"imoveis/models"
	"log"
)

func AuthUsuarios() map[string]string {
	var usuarios []models.Usuario
	databases.DB.Find(&usuarios)
	if len(usuarios) == 0 {
		usuario := models.Usuario{
			Usuario: "admin",
			Senha:   base64.StdEncoding.EncodeToString([]byte("admin")),
			Tipo:    "administrador",
		}
		databases.DB.Save(&usuario)
	}
	logins := map[string]string{}
	for _, u := range usuarios {
		logins[u.Usuario] = func() string {
			s, err := base64.StdEncoding.DecodeString(u.Senha)
			if err != nil {
				log.Panic("Não conseguimos acessar as senhas:", err)
			}
			return string(s)
		}()
	}
	return logins
}
