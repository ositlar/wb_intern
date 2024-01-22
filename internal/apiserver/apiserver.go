package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/ositlar/wb_intern_1/internal/store/sqlstore"
)

func Start(config *Config) error {
	db, err := connectToDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	store := sqlstore.NewStore(db)
	srv := newServer(*store)
	defer srv.sc.Close()
	return http.ListenAndServe(":8080", srv)
}

func connectToDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
