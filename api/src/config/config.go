package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DbData struct {
	Host        string
	Port        string
	DbName      string
	RolName     string
	RolPassword string
}

type Config struct {
	//Dados para conexao com a base de dados
	DbData DbData
	//Porta onde a API vai rodar
	APIPort string
	//Chave que vai ser usada para assinar o token
	SecretKey []byte
}

var (
	GlobalConfig Config
)

// Executa e carrega as variáveis de ambiente
func init() {
	LoadEnv()
}

// Inicializa as variáveis de ambiente
func LoadEnv() Config {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	GlobalConfig = Config{
		DbData: DbData{
			Host:        os.Getenv("DB_HOST"),
			Port:        os.Getenv("DB_PORT"),
			DbName:      os.Getenv("DB_NAME"),
			RolName:     os.Getenv("ROL_NAME"),
			RolPassword: os.Getenv("ROL_PASSWORD"),
		},
		APIPort:   os.Getenv("API_PORT"),
		SecretKey: []byte(os.Getenv("SECRET_KEY")),
	}

	return GlobalConfig
}
