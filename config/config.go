package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	DBHOST string
	DBPORT string
	DBUSER string
	DBPASS string
	DBNAME string
}

func Load() *Config{
	err:=godotenv.Load()
	if err!=nil{
		log.Println("No env variable detected")
	}

	cfg:=&Config{
		DBHOST: os.Getenv("DB_HOST"),
		DBPORT: os.Getenv("DB_PORT"),
		DBUSER: os.Getenv("DB_USER"),
		DBPASS: os.Getenv("DB_PASS"),
		DBNAME: os.Getenv("DB_NAME"),
	}

	if cfg.DBHOST=="" || cfg.DBPORT=="" || cfg.DBUSER=="" || cfg.DBPASS=="" || cfg.DBNAME=="" {
		log.Fatal("Database env var are not properly set")
	}

	return cfg
}