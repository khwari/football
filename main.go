package main

import (
	"net/http"
	"github.com/subutai-io/agent/log"
	"football/news"
)

var (
	srv      *http.Server
)
func main()  {
	http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/news", news.GetNews)

	log.Info("Server has started...")

	srv = &http.Server{
		Addr:    ":" + "8080",
		Handler: nil,
	}
	srv.ListenAndServe()

}
//d
func welcome(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome"))
}

