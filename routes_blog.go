package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type jsonPosts struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func mapPostsToJson(p *Posts) jsonPosts {
	return jsonPosts{
		ID:      p.ID,
		Title:   p.Title,
		Content: p.Content,
		Author:  p.Author,
	}
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello World")
	}
}

func (s *server) handlePostsList() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		posts, err := s.blog.GetPosts()

		if err != nil {
			log.Printf("Ne peut lire les posts. err=%v\n", err)
			s.respond(rw, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]jsonPosts, len(posts))
		for i, m := range posts {
			resp[i] = mapPostsToJson(m)
		}

		s.respond(rw, r, resp, http.StatusOK)
	}
}

func (s *server) handleOnePost() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			log.Printf("Ne peut parse l'id en int. err=%v\n", err)
			s.respond(rw, r, nil, http.StatusBadRequest)
			return
		}

		p, err := s.blog.GetPostById(id)
		if err != nil {
			log.Printf("Ne peut lire le post. err=%v\n", err)
			s.respond(rw, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = mapPostsToJson(p)
		s.respond(rw, r, resp, http.StatusOK)
	}
}
