package main

import (
	"imoveis/controllers/sistemaVelho"
	"imoveis/databases"
	"imoveis/routers"
	"os"
)

func init() {
	os.Setenv("TZ", "America/Sao_Paulo")
}

func main() {
	databases.ConectaDB()
	sistemaVelho.MigraAlugueis()
	routers.IniciaRoteamento()
}
