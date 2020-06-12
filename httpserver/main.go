package main

import (
	"fmt"
	"net/http"
)

func main() {
	// r.URL.Query().Get("token") でGETパラメータを取得できる
	// r.FormValue("email") 等でPOSTパラメータを取得できる
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	// JS, CSS, ImageなどのStaticコンテンツを http.FileServerでPATHを指定
	// PATHは http.Dir を使用
	fs := http.FileServer(http.Dir("static/"))

	// URLを指定
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// ポートをリッスン
	http.ListenAndServe(":80", nil)
}
