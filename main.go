package main

import (
	"fmt"
	"net/http"
)

func main() {

	// 顯示檔案
	http.Handle("/", http.FileServer(http.Dir("./public")))

	// 自定義功能
	http.HandleFunc("/shopee", func(w http.ResponseWriter, r *http.Request) {
		// r 是從網頁讀取
		// w 是寫到網頁

		query := r.URL.Query()
		buyWhat := query.Get("buy")
		buyCount := query.Get("count")
		fmt.Println("buy => ", buyWhat)
		fmt.Println("count => ", buyCount)

		var name string
		name = "李家家"
		w.Write([]byte(name))
	})

	// 開始伺服器
	err := http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		fmt.Println("開始伺服器 有錯誤 =>", err)
	}
}
