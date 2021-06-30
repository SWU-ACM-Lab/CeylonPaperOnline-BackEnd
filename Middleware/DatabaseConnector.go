package Middleware

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type dbConnection struct {
	Db *sql.DB
	connected bool
}

func (d* dbConnection) Connect (config DatabaseConfig) bool {
	if !d.connected {
		path := config.GeneratePath()
		db, err := sql.Open("mysql", path)
		if err != nil {
			Console.Log(err, "Generate Database Path")
			d.connected = false
		} else {
			err = db.Ping()
			if err != nil {
				Console.Log(err, "Connect to Database")
				d.connected = false
			} else {
				d.Db = db
				d.connected = true
			}
		}
	}

	return d.connected
}

type QueryConsole struct {
	dbConnection
}