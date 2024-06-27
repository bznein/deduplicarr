package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

var db *sql.DB

type column struct {
	name       string
	dbType     string
	nullable   bool
	primaryKey bool
}

func NewColumn(name string, dbType string, null bool, primary bool) column {
	return column{name: name, dbType: dbType, nullable: null, primaryKey: primary}
}

// TODO not foo.db xD
func EnsureDBConnection() (err error) {
	if db != nil {
		log.Warn("tried to initialize DB more than once")
		return nil
	}
	db, err = sql.Open("sqlite3", "./foo.db")

	return err
}

func CloseDB() {
	db.Close()
}

func EnsureTableExists(name string, columns ...column) error {
	sqlStmt := fmt.Sprintf("create table if not exists %s (", name)

	for i, c := range columns {
		sqlStmt = fmt.Sprintf("%s %s %s", sqlStmt, c.name, c.dbType)
		if !c.nullable {
			sqlStmt += " not null"
		}
		if c.primaryKey {
			sqlStmt += " primary key"
		}
		if i < len(columns)-1 {
			sqlStmt += ", "
		}
	}
	sqlStmt += ");"

	log.Debugf("Executing SQL statement %s", sqlStmt)
	_, err := db.Exec(sqlStmt)

	return err
}
