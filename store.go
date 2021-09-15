package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// ! Interface de nos functions
type Blog interface {
	Open() error
	Close() error
}

type dbBlog struct {
	db *sqlx.DB
}

// ! Func Open() de notre BDD
func (blog *dbBlog) Open() error {
	// ? Connexion à la DB avec sqlite3 comme pilote
	db, err := sqlx.Connect("sqlite3", "goblog")
	if err != nil {
		return err
	}

	fmt.Printf("Connexion à la base de donnée effectué!")

	blog.db = db
	return nil
}

// ! Func Close() de notre BDD
func (blog *dbBlog) Close() error {
	return blog.db.Close()
}
