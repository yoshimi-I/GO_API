package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

func HelloHandlers(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello World!")
}

func PostAriticleHandlers(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article")
}

func ArticleListsHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	var err error
	if p, ok := queryMap["page"]; ok && len(p) > 0 { //okは受け取れた場合tureがはいる
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "クエリがおかしい", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}
	resString := fmt.Sprintf("記事番号は%dです", page)
	io.WriteString(w, resString)
}

func ArticleDetalsHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "パスパラメータがおかしい", http.StatusBadRequest)
	}
	resString := fmt.Sprintf("記事の番号は%dです", articleID)
	io.WriteString(w, resString)
}

func PostNiceHnadler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice")
}
func PostCommentHnadler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment")
}
