package main

import (
	"github.com/gorilla/mux"
	"github.com/yoshimi-I/GO_API/handlers"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handlers.HelloHandlers).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostAriticleHandlers)
	r.HandleFunc("/article/list", handlers.ArticleListsHandler)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetalsHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHnadler)
	r.HandleFunc("/article/comment", handlers.PostCommentHnadler)

	log.Println("サーバーを8080で立ち上げています")
	log.Fatal(http.ListenAndServe(":8080", r))
}
