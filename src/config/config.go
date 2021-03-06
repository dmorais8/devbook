package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// URL de conexão com o banco de dados.
	StringConexaoBanco = ""

	// Porta onde a API vai escutar.
	Porta = 0

	// Token Secret Key
	SecretKey []byte
)

func Carregar() {

	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

}
