package main

import (
	"imoveis/databases"
	"imoveis/routers"
	"os"
)

func init() {
	os.Setenv("TZ", "America/Sao_Paulo")
}

func main() {
	databases.ConectaDB()
	routers.IniciaRoteamento()
}
