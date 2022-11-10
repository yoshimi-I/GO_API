# Go言語のAPIチュートリアル
# 目次
1. [ハンドラーの作成](#anchor1)
2. [HTTPメソッド](#anchor2)
3. [ステータスコード](#anchor3)
4. [パスパラメータ](#anchor4)
5. [クエリパラメータ](#anchor5)


# 1. ハンドラーの作成 <a id="anchor1"></a>
- Goは簡単にローカルを立ち上げてハンドラーを作れる
- ハンドラーとは簡単にいうと、urlのリクエストに応じたレスポンスを返すものである
```go
func HelloHandlers(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Hello World!")
	}
http.HandleFunc("/hello", handlers.HelloHandlers)
log.Println("サーバーを8080で立ち上げています")
log.Fatal(http.ListenAndServe(":8080", nil))
```
- 基本的にはhttp.Responsewrierとhttp.Requestを引数に取る
  - http.HandleFuncにてパスとその時呼び出す関数を記載する
  - http.Requestがポインタ値の理由はその中のHTTPメソッドを参照するためである。
- 基本的な流れは,req *http.Requestでメソッドを分類しResponseWriterを引数にとって処理を書く

# 2. HTTPメソッド <a id="anchor2"></a>
- 基本的には以下の通り
- | メソッド名 | 意味                         |
  | ---------- | ---------------------------- |
  | GET        | データの取得                 |
  | POST       | データをクライアントから送信 |
  | PUT        | データの更新                 |
  | DELETE     | データの削除                 |

- メソッドはreq.Methodの中に入っている
    - これをもとにどのメソッドなのかを判断する
# 3. ステータスコード <a id="anchor3"></a>
- 以下にコードの説明を示す
- | ステータスコード | テキスト             | 意味                         |
  | ---------------- | -------------------- | ---------------------------- |
  | 200              | OK                   | 成功                         |
  | 400              | Bad request          | リクエストの値が不正         |
  | 403              | Forbidden            | ユーザのアクセス権限がない   |
  | 404              | Not found            | URLが存在しない              |
  | 405              | Method not allowd    | HTTPメソッドの許可が降りない |
  | 500              | Internal Servererror | サーバー内部でのエラー       | 
- これをGoで使う場合はhttpパッケージを使う
```go
https.StatusOK　              = 200
https.StatusBadREquest        = 400
https.StatusForbidden         = 403
https.StatusNotFound          = 404
https.StatusMethodNotAllowd   = 405
```
- といった感じ
### 具体例
```go
func HelloHandlers(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Hello World!")
	}else {
		http.Error(w,"ざまあ,405です",http.StatusMethodNotAllowed)
	}
```
# 3. パスパラメータ <a id="anchor4"></a>
- 簡単にいうとurlの後ろにつける変数
  - hello/1とかhello/2とか全部ハンドラーを描くのはめんどくさいから変数にしたいよね
  - このULRの中にあるIDをパラメータとして変数かしてハンドラの中で利用することを言う
  ### main.go
- ```go
  import (
	"github.com/gorilla/mux"
  )
  r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetalsHandler).Methods(http.MethodGet)
  ```
  ### handler.go
  ```go
    func ArticleDetalsHandler(w http.ResponseWriter, req *http.Request) {
	articleID,err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil{
		http.Error(w,"パスパラメータがおかしい",http.StatusBadRequest)
	}
	resString := fmt.Sprintf("記事の番号は%dです", articleID)
	io.WriteString(w, resString)
  }
  ```

- idをstrconv.Atoi(mux.Vars(req)["id"])で受け取ってエラーハンドリングをする
- Vars関数はmuxに実装されている関数であり、引数に辞書型をとってくれるため、そのkeyを["id"]のように指定することで取得可能
- これをAtoiを用いてint型にキャスト
# 4. クエリパラメータ <a id="anchor5"></a>
- 簡単に言うとパラメータの後ろに条件を付け加えるやつ
  - 条件一致の検索とかするあれです。
  - http://gaforum.jp/?s=gaiqみたいなやつです。
  - 