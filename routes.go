package main

func (s *server) routes() {
	s.router.HandleFunc("/", s.handleIndex()).Methods("GET")
	s.router.HandleFunc("/api/posts/{id:[0-9]+}", s.handleOnePost()).Methods("GET")
	s.router.HandleFunc("/api/posts/", s.handlePostsList()).Methods("GET")
	s.router.HandleFunc("/api/posts/", s.handleCreatePost()).Methods("POST")
}
