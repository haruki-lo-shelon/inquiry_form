package req_handler // 独自のHTTPリクエストハンドラパッケージ

import (
    "fmt"
    "html/template"
    "net/http"
)

// 入力内容の確認画面
func HandlerConfirm(w http.ResponseWriter, req *http.Request) {
    // テンプレートをパースする
    tpl := template.Must(template.ParseFiles("templates/confirm.gtpl"))

    // テンプレートに出力する値をマップにセット
    values := map[string]string{
        "name": req.FormValue("name"),
        "email":    req.FormValue("email"),
        "inquiry":  req.FormValue("inquiry"),
				"hid_name": req.FormValue("name"),
        "hid_email":    req.FormValue("email"),
        "hid_inquiry":  req.FormValue("inquiry"),
    }

    // マップを展開してテンプレートを出力する
    if err := tpl.ExecuteTemplate(w, "confirm.gtpl", values); err != nil {
        fmt.Println(err)
    }
}
