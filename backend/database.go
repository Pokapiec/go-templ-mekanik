package backend

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT,
	password TEXT,
	app_role TEXT
);

CREATE TABLE IF NOT EXISTS assignment (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	assignment_name TEXT,
	assignment_type INTEGER,
	assigned_to INTEGER REFERENCES users(id),
	assigned_by INTEGER REFERENCES users(id),
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);	
`

func ConnectDB() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", "./database.db")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)
	// db.MustExec(`INSERT INTO users (username, password, app_role) VALUES ("admin", "admin", "admin")`)
	return db
}
