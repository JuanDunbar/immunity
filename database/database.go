package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/juandunbar/immunity/config"
)

type Database interface {
	Connect(c *config.Config) error
	Disconnect() error
	Query(query string) error
	Execute(query string) error
}

func NewDatabase() Database {
	return new(postgres)
}

type postgres struct {
	DB *sql.DB
}

func (pg *postgres) Connect(c *config.Config) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Database.Host, c.Database.Port, c.Database.User,
		c.Database.Password, c.Database.DBname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	pg.DB = db

	return nil
}

func (pg *postgres) Disconnect() error {
	return pg.DB.Close()
}

func (pg *postgres) Query(query string) error {
	return nil
}

func (pg *postgres) Execute(query string) error {
	return nil
}
