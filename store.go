package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// ! Interface de nos functions
type Blog interface {
	Open() error
	Close() error

	GetPosts() ([]*Posts, error)
}

type dbBlog struct {
	db *sqlx.DB
}

var schema = `
CREATE TABLE IF NOT EXISTS blog
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT,
	content TEXT
	autho TEXT
)
`

// ! Func Open() de notre BDD
func (blog *dbBlog) Open() error {
	// ? Connexion à la DB avec sqlite3 comme pilote
	db, err := sqlx.Connect("sqlite3", "goblog.db")
	if err != nil {
		return err
	}

	fmt.Printf("Connexion à la base de donnée effectué!\n")

	db.MustExec(schema)
	blog.db = db
	return nil
}

// ! Func Close() de notre BDD
func (blog *dbBlog) Close() error {
	return blog.db.Close()
}
