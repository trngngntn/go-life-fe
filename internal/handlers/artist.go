package handlers

import (
	"net/http"

	"go-life-fe/internal/components"
	"go-life-fe/internal/entities"
	"go-life-fe/internal/utils"

	"github.com/gorilla/mux"
)

func InitArtistHandler(subRouter *mux.Router) {
	subRouter.HandleFunc("", artistGETHandler).Methods("GET")
	subRouter.HandleFunc("/{id:[0-9]+}", artistGETHandler).Methods("GET")
	subRouter.HandleFunc("", artistPOSTHandler).Methods("POST")
	subRouter.HandleFunc("/{id:[0-9]+}", artistPUTHandler).Methods("PUT")
	subRouter.HandleFunc("/{id:[0-9]+}", artistDELETEHandler).Methods("DELETE")
}

func artistGETHandler(wr http.ResponseWriter, req *http.Request) {
	if ID := utils.GetID(req); ID != -1 {
		artist := entities.GetArtistByID(ID, utils.GetCookie(req, "token"))
		components.RenderPage("artist", wr, artist, 1)
	}
}

func artistPOSTHandler(writer http.ResponseWriter, req *http.Request) {

}

func artistPUTHandler(writer http.ResponseWriter, req *http.Request) {

}

func artistDELETEHandler(writer http.ResponseWriter, req *http.Request) {

}
