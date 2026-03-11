package db

import (
	"database/sql"
	"fmt"
	"log"
	"mini_search_engine/config"
	_ "github.com/lib/pq"
)

func Initializer(cfg *config.Config) *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHOST, cfg.DBPORT, cfg.DBUSER, cfg.DBPASS, cfg.DBNAME,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect the DB: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping the postgres: %v", err)
	}

	log.Println("postgres connected succesfully")
	return db
}
