package main

import (

	// "fmt"
	// "html/template"
	"net/http"
	"inquiry/req_handler"

	// "github.com/gin-gonic/gin"
)

func main() {

		// "user-form"へのリクエストを関数で処理する
    http.HandleFunc("/form", req_handler.HandlerForm)

    // "user-confirm"へのリクエストを関数で処理する
    http.HandleFunc("/confirm", req_handler.HandlerConfirm)

    // "user-registered"へのリクエストを関数で処理する
    http.HandleFunc("/send", req_handler.HandlerSend)

		// css・js・イメージファイル等の静的ファイル格納パス
    http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("asset/"))))

		// サーバーを起動
    http.ListenAndServe(":8080", nil)
    // err := http.ListenAndServeTLS(":10443", "crt/localhost.crt", "crt/localhost.key", nil)
    // if err != nil {
    //     fmt.Printf("ERROR : %s", err)
    // }
}
