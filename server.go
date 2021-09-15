package main

type server struct {
	blog Blog
}

func newServer() *server {
	s := &server{}
	return s
}
