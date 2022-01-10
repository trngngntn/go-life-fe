package main

import "go-life-fe/cmd/life"

func main() {
	life.Init()
	life.Start()
}

// func handler(writer http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(writer, "Hey hey, my my")
// }

// func viewHandler(writer http.ResponseWriter, req *http.Request) {
// 	title := req.URL.Path[len("/view/"):] // from after view/ to end
// 	page, err := loadPage(title)
// 	if err != nil {
// 		page = &Page{Title: title}
// 	}
// 	template, _ := template.ParseFiles("web/frame.html")
// 	template.Execute(writer, page)
// }
