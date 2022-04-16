// Package datasource implements datasource connection.
package datasource

import (
	"database/sql"
	"fmt"
	"github.com/hexhoc/go-mall-api/config"
	_ "github.com/lib/pq"
)

// NewPostgresConnection New -.
func NewPostgresConnection(datasource *config.Datasource) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		datasource.Username,
		datasource.Password,
		datasource.Host,
		datasource.Port,
		datasource.Database,
		datasource.Sslmode)

	db, err := sql.Open("postgres", connStr)

	return db, err
}
