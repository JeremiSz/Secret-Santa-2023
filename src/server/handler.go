package server

import (
	"log"
	"net/http"
)

type Server struct {
	http.Server
}

func NewServer() *Server {
	router := http.NewServeMux()
	router.HandleFunc("/", staticViewProvider)

	server := &Server{http.Server{
		Addr:    ":8080",
		Handler: router,
	}}
	return server
}
func staticViewProvider(w http.ResponseWriter, r *http.Request) {
	log.Println("serving" + r.URL.Path)
	switch r.URL.Path {
	case "/snow.css":
		http.ServeFile(w, r, "./site/css/snow.css")
	default:
		http.ServeFile(w, r, "./site/html/index.html")
	}
}
