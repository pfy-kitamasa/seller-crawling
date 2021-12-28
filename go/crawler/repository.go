package crawler

import (
	"database/sql"
	"fmt"

	"github.com/panforyou/seller-finding/go/crawler/repository"

	// cloudsqlproxy "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	_ "github.com/lib/pq"
)

func NewDB(host string, port int, dbname, user, password string, env string) (*DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		host, port, dbname, user, password)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	queries := repository.New(db)

	return &DB{
		db,
		queries,
	}, nil
}

type DB struct {
	*sql.DB
	*repository.Queries
}

func (d *DB) Close() error {
	return d.DB.Close()
}
