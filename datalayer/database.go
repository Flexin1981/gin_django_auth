package datalayer

import (
	"os"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)


var (
	databaseSelections = map[string]func () (*sql.DB, schema.Dialect){
		"postgres": getPostgress,
	}
)

func getPostgress() (*sql.DB, schema.Dialect) {
	return sql.OpenDB(
		pgdriver.NewConnector(
			pgdriver.WithDSN(os.Getenv(DatabaseConnectionEnvironmentVariable)),
		),
	),
	pgdialect.New()
}


func GetDatabaseConnection() *bun.DB {
	dataBase, dialect := databaseSelections[os.Getenv(DatabaseDialect)]()
	return bun.NewDB(dataBase, dialect)
}