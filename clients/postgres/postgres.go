package postgres

import (
	"context"
	"errors"
	"log"

	"github.com/baturalpk/photo-bucket/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

func GetConnectionString() (connstring string, err error) {
	switch config.Get().Env {
	case "prod":
		connstring = config.Get().PostgresURI.Prod
	case "dev":
		connstring = config.Get().PostgresURI.Dev
	case "test":
		connstring = config.Get().PostgresURI.Test
	default:
		err = errors.New("required config key('ENV') is not defined")
	}
	return
}

func InitConnection() error {
	connstring, err := GetConnectionString()
	if err != nil {
		return err
	}

	log.Println("trying to establish postgres connection...")
	pool, err := pgxpool.Connect(context.Background(), connstring)
	if err != nil {
		return err
	}

	if err = pool.Ping(context.Background()); err != nil {
		return err
	}
	db = pool
	log.Println("postgres connection pool is created successfully")
	return nil
}

func CloseConnections() {
	db.Close()
}

func DB() *pgxpool.Pool {
	return db
}
