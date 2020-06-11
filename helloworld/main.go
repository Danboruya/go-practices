package main

import (
	"fmt"
	// net/http packageにはHTTPクライアントとHTTPサーバーが含まれている
	"net/http"
)

func main() {
	// ブラウザ,APIリクエスト,HTTPクライアント等からの全HTTP接続を受け付けるハンドラ
	// http.ResponseWriter : text/htmlレスポンスを書き込む
	// http.Request : URL, header fieldなどの全HTTPリクエスト情報が格納されている
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	// Requestリスナー, 指定したポートをlistenする
	// 今回はHTTPのため80番
	http.ListenAndServe(":80", nil)
}
