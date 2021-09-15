package main

import (
	"fmt"
	"log"
	"net/http"
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
