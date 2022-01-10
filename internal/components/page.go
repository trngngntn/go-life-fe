package components

import (
	"bytes"
	"fmt"
	htmltmpl "html/template"
	"io"
	"log"
	"os"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

func loadPLayerFrame() {

}

func loadManagementFrame() {

}

func loadPage(name string) (*Page, error) {
	filename := "web/" + name + ".html"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: name, Body: body}, nil
}

func RenderPage(name string, wr io.Writer, data interface{}, frameType int) {

	funcMap := htmltmpl.FuncMap{
		"time2txt": func(t int) string {
			return fmt.Sprintf("%d:%02d", t/60, t%60)
		},
	}

	tmpl, err := htmltmpl.New(name + ".gohtml").Funcs(funcMap).ParseFiles("web/" + name + ".gohtml")
	if err != nil {
		log.Fatal(err)
	}

	if frameType == 0 {
		err = tmpl.Execute(wr, data)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		var buffer bytes.Buffer
		err = tmpl.Execute(&buffer, data)
		if err != nil {
			log.Fatal(err)
		}
		frameTempl, err := template.ParseFiles("web/frame.html")
		if err != nil {
			log.Fatal(err)
		}
		frameTempl.Execute(wr, buffer.String())
	}
}
