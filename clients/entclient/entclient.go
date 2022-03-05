package entclient

import (
	"context"
	"database/sql"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/baturalpk/photo-bucket/clients/postgres"
	"github.com/baturalpk/photo-bucket/ent"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var client *ent.Client

func InitConnection() error {
	connstring, err := postgres.GetConnectionString()
	if err != nil {
		return err
	}

	db, err := sql.Open("pgx", connstring)
	if err != nil {
		return err
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	client = ent.NewClient(ent.Driver(drv))

	if err := client.Schema.Create(context.Background()); err != nil {
		return err
	}
	return nil
}

func CloseConnection() {
	_ = client.Close()
}

func Client() *ent.Client {
	return client
}
