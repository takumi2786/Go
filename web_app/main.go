// フレームワークなしでの，ウェブアプリケーションの試作
// 参考：https://golang.org/doc/articles/wiki/

package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// Pageという構造体を作成
type Page struct {
	Title string
	Body  []byte
}

// データをtxtファイルに保存する関数．
// 0600はPORT？
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// p:テンプレートに渡す構造体．
	// w:http.ResponseWriter ここに何かを書くと，ページに反映される．
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):] //スライス的な使い方？
	// println(title)
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body) //htmlのページとして表示される．// w変数に情報を書き込んでいる．
	renderTemplate(w, "view", p) // これは，テンプレートファイルをコンパイルして，wに情報を書き込んでいる．
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):] //スライス的な使い方？
	// println(title)
	p, err := loadPage(title)
	if err != nil { //editの場合は，ページファイルが存在しなくても，新規で作成
		p = &Page{Title: title}
	}
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body) //htmlのページとして表示される．// w変数に情報を書き込んでいる．
	renderTemplate(w, "edit", p) // これは，テンプレートファイルをコンパイルして，wに情報を書き込んでいる．
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):] //スライス的な使い方？
	body := r.FormValue("body")         //name==bodyのフォーム情報を取得
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// リダイレクト，つまりはページ移動．
	// ３つめの引数は何？
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// ウェブサーバの立ち上げ
func main() {
	// /view/配下のルーティングに相当する．
	// http://localhost:8080/view/へのアクセスがあった場合に，viewHandlerを実行する．
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	// http.ListenAndServe: :8080が生きている限り，サーバを立ち上げ続ける？
	// 第一引数：ポート
	// 第二引数：ハンドラ
	// log.Fatal：errorがあった場合に，ログとして表示する関数．
	log.Fatal(http.ListenAndServe(":8080", nil))

}
