package main

import "fmt"

type Posts struct {
	ID      int64  `db:"id"`
	Title   string `db:"title"`
	Content string `db:"content"`
	Author  string `db:"author"`
}

func (p Posts) String() string {
	return fmt.Sprintf("id=%v, title=%v, content=%v, author=%v\n", p.ID, p.Title, p.Content, p.Author)
}
