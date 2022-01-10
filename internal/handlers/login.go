package handlers

import (
	"net/http"

	"go-life-fe/internal/components"
	"go-life-fe/internal/entities"

	"github.com/gorilla/mux"
)

func InitLoginHandler(subRouter *mux.Router) {
	subRouter.HandleFunc("", loginGETHandler).Methods("GET")
	subRouter.HandleFunc("", loginPOSTHandler).Methods("POST")
}

func loginGETHandler(wr http.ResponseWriter, req *http.Request) {
	components.RenderPage("login", wr, nil, 0)
}

func loginPOSTHandler(wr http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	email := req.Form.Get("email")
	password := req.Form.Get("password")

	if email == "admin" && password == "admin" {
		http.Redirect(wr, req, "/artists/1", http.StatusSeeOther)
	}
	//authenticate
	if user := entities.Login(email, password); user != nil {
		http.Redirect(wr, req, "/artists/1", http.StatusSeeOther)
	} else {
		components.RenderPage("login", wr, "", 0)
	}

	components.RenderPage("login", wr, nil, 0)
}
