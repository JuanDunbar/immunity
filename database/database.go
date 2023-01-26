package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// TODO move to config
const (
	host     = "localhost"
	port     = 5432
	user     = ""
	password = ""
	dbname   = "immunity"
)

var dbClient *sql.DB

func Connect(ctx context.Context) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	dbClient = db
	return nil
}

func Disconnect() error {
	return dbClient.Close()
}

func GetClient() *sql.DB {
	return dbClient
}
