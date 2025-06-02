package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Porta              = 0
	SecretKey          []byte
)

// Carregar inicializes the ambient variables
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	Porta, erro = strconv.Atoi(os.Getenv("PORT"))
	if erro != nil {
		Porta = 9000
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local&tls=true",
		os.Getenv("MYSQLUSER"),
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		os.Getenv("MYSQLHOST"),
		os.Getenv("MYSQLDATABASE"),
	)
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
