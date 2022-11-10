package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/yoshimi-I/GO_API/handlers"
	"log"
	"net/http"
	"time"
)

type Commet struct {
	CommentID int       `json:"comment_id"`
	ArticleID int       `json:"article_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	UserName    string    `json:"user_name"`
	NiceNum     int       `json:"nice_num"`
	CommentList []Commet  `json:"comment_list"`
	CreatedAt   time.Time `json:"created_at"`
}

func main() {
	comment1 := Commet{
		CommentID: 1,
		ArticleID: 1,
		Message:   "ひん",
		CreatedAt: time.Time{},
	}
	comment2 := Commet{
		CommentID: 2,
		ArticleID: 1,
		Message:   "加藤純一最強",
		CreatedAt: time.Time{},
	}
	article1 := Article{
		ID:          1,
		Title:       "衛門",
		Contents:    "加藤純一最強",
		UserName:    "品川彩人",
		NiceNum:     1,
		CommentList: []Commet{comment1, comment2},
		CreatedAt:   time.Time{},
	}
	jsonData, err := json.Marshal(article1)
	if err != nil {
		fmt.Printf("josnの変換に失敗しまいした:%s", err)
	}
	fmt.Printf("%s", jsonData)

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
