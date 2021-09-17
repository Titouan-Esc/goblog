package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Sprintf(os.Stderr.Name(), "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	srv := newServer()
	srv.blog = &dbBlog{}
	err := srv.blog.Open()
	if err != nil {
		return err
	}
	defer srv.blog.Close()

	http.HandleFunc("/", srv.serveHTTP)
	log.Printf("Serving HTTP on, PORT 8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}
	return nil
}
