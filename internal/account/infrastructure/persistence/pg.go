package persistence

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Repositories struct {
	db *sql.DB
}

func NewRepositories(user string, password string, host string, port string, dbName string) (*Repositories, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	db, openErr := sql.Open("postgres", connectionString)
	if openErr != nil {
		return nil, openErr
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	return &Repositories{
		db: db,
	}, nil
}

func (r *Repositories) Close() error {
	return r.db.Close()
}
