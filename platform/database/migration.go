package database

import (
	"embed"
	"fmt"
	"log"
	"movie_storage/pkg/configs"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var sqlFiles embed.FS

func (db *Database) MigrateDB(cfg *configs.ConfigPostgressDB) error {
	ok := db.CheckAndMigrateTables()
	if ok {
		fmt.Printf("вернул нил")
		return nil
	}
	d, err := iofs.New(sqlFiles, "migrations")
	if err != nil {
		log.Fatal(err)
	}
    dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
    
	m, err := migrate.NewWithSourceInstance("migration_embeded_sql_files", d, dsn)
	if err != nil {
		log.Fatal(err)
	}

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Println(err)
        return err
    }

    return nil
}


func (db *Database) CheckAndMigrateTables() bool {

	tables := []string{"users", "actors", "movies"}


	for _, table := range tables {
		if !db.checkTableExists(table) {
			return false
	}
}
return true
}


func (db *Database) checkTableExists(tableName string) bool {
	query := "SELECT EXISTS (SELECT FROM pg_tables WHERE schemaname = 'public' AND tablename = $1);"
	var exists bool
	err := db.db.QueryRow(query, tableName).Scan(&exists)
	if err != nil {
	}
	return exists
}