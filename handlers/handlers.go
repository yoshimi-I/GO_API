package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func HelloHandlers(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Hello World!")
	}

}

func PostAriticleHandlers(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article")
}

func ArticleListsHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List")
}

func ArticleDetalsHandler(w http.ResponseWriter, req *http.Request) {
	articleID := 1
	resString := fmt.Sprintf("記事の番号は%dです", articleID)
	io.WriteString(w, resString)
}

func PostNiceHnadler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice")
}
func PostCommentHnadler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment")
}
