package databases

import (
	"fmt"
	"imoveis/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaDB() {
	user := os.Getenv("DBUSER")
	pass := os.Getenv("DBPASS")
	host := os.Getenv("DBHOST")
	database := os.Getenv("DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, database)
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal("erro ao conectar ao banco de dados", err)
	}
	DB.AutoMigrate(&models.Cliente{})
	DB.AutoMigrate(&models.Imovel{})
	DB.AutoMigrate(&models.Usuario{})
	DB.AutoMigrate(&models.Energia{})
}
