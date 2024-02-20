package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/juandunbar/immunity/config"
)

var schema = `
-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS rules_id_seq;

-- Table Definition
CREATE TABLE "public"."rules" (
    "id" int4 NOT NULL DEFAULT nextval('rules_id_seq'::regclass),
    "query" varchar NOT NULL,
    "description" varchar NOT NULL,
    "action" varchar NOT NULL,
    "last_used" timestamp NOT NULL,
    "disabled" bool NOT NULL DEFAULT false,
    PRIMARY KEY ("id")
);
`

type Database interface {
	Connect(c *config.Config) error
	Disconnect() error
	Query(object any, query string) error
	Execute(query string, args any) error
}

func NewDatabase() Database {
	return new(postgres)
}

type postgres struct {
	DB *sqlx.DB
}

func (pg *postgres) Connect(c *config.Config) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Database.Host, c.Database.Port, c.Database.User,
		c.Database.Password, c.Database.DBname)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return err
	}
	//db.MustExec(schema)
	pg.DB = db

	return nil
}

func (pg *postgres) Disconnect() error {
	return pg.DB.Close()
}

func (pg *postgres) Query(object any, query string) error {
	err := pg.DB.Select(object, query)
	if err != nil {
		return err
	}
	return nil
}

func (pg *postgres) Execute(query string, arg any) error {
	_, err := pg.DB.NamedExec(query, arg)
	return err
}
