package main

import (
	"log"
	"net/http"
	"yoshimi-I/GO_API/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandlers)
	http.HandleFunc("/article", handlers.PostAriticleHandlers)
	http.HandleFunc("/article/list", handlers.ArticleListsHandler)
	http.HandleFunc("/article/1", handlers.ArticleDetalsHandler)
	http.HandleFunc("/article/nice", handlers.PostNiceHnadler)
	http.HandleFunc("/article/comment", handlers.PostCommentHnadler)

	log.Println("サーバーを8080で立ち上げています")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
