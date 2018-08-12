package main

import (
	"html/template"
	"net/http"
)

type Person struct {
	Name string // 命名要大寫開頭
}

func indexHandle(w http.ResponseWriter, r *http.Request) {
	t := template.New("some template") //创建一个模板
	t, _ = t.ParseFiles("index.html")  // 讀取 HTML 檔案
	p := Person{Name: "Loli"}          // 要送到前端的訊息
	t.Execute(w, p)

	//t := template.New("some template") //创建一个模板
	//t, _ = t.ParseFiles("tmpl/welcome.html", nil)  //解析模板文件
	//user := GetUser() //获取当前用户信息
	//t.Execute(w, user)  //执行模板的merger操作

}

func main() {
	http.HandleFunc("/", indexHandle) // 當瀏覽器輸入根目錄時會呼叫 indexHandle() 涵式
	http.ListenAndServe(":80", nil)
}
