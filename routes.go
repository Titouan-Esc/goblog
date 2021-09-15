package main

func (s *server) routes() {
	s.router.HandleFunc("/", s.handleIndex()).Methods("GET")
	s.router.HandleFunc("/api/posts/", s.handlePostsList()).Methods("GET")
}
