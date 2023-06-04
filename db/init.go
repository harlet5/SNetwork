package db

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rubenv/sql-migrate"
)

var Db *sql.DB

func InitDb() error {
	var err error
	Db, err = sql.Open("sqlite3", "./db/forum.db")
	return err
}

func Migrate() error {

	// flag.Parse()
	// if *migratePtr == true {
		migrations := &migrate.FileMigrationSource{
			Dir: "db/migrations/default",
		}
	
		n, err := migrate.Exec(Db, "sqlite3", migrations, migrate.Up)
		if err != nil {
			log.Println(err)
			return err
		}
		log.Printf("Applied %d migrations!\n", n)
	// }
	return nil

}
