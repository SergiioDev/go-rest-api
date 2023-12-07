package config

import (
	"log"

	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

type Conf struct {
	DB     ConfDB
	Server ConfServer
}

type ConfServer struct {
	Port string `env:"SERVER_PORT,required"`
}

type ConfDB struct {
	Host     string `env:"DB_HOST,required"`
	Port     uint16 `env:"DB_PORT,required"`
	UserName string `env:"DB_USER,required"`
	Password string `env:"DB_PASS,required"`
	DBName   string `env:"DB_NAME,required"`
	Debug    bool   `env:"DB_DEBUG,required"`
}

func New() *Conf {
	loadEnvFile()
	var conf Conf
	err := envdecode.StrictDecode(&conf)
	if err != nil {
		log.Fatalf("Error decoding env variables into conf struct: %s", err.Error())
	}
	return &conf
}

func NewDB() *ConfDB {
	loadEnvFile()
	var confDb ConfDB
	err := envdecode.Decode(&confDb)
	if err != nil {
		log.Fatalf("Error decoding env variables into confDb struct %s", err.Error())
	}
	return &confDb
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}
}
