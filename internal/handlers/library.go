package handlers

import (
	"net/http"

	"go-life-fe/internal/components"

	"github.com/gorilla/mux"
)

func InitLibraryHandler(subRouter *mux.Router) {
	subRouter.HandleFunc("", libraryGETHandler).Methods("GET")
	subRouter.HandleFunc("/{id:[0-9]+}", libraryGETHandler).Methods("GET")
	subRouter.HandleFunc("", libraryPOSTHandler).Methods("POST")
	subRouter.HandleFunc("/{id:[0-9]+}", libraryPUTHandler).Methods("PUT")
	subRouter.HandleFunc("/{id:[0-9]+}", libraryDELETEHandler).Methods("DELETE")
}

func libraryGETHandler(wr http.ResponseWriter, req *http.Request) {
	components.RenderPage("library", wr, nil, 1)
}

func libraryPOSTHandler(writer http.ResponseWriter, req *http.Request) {

}

func libraryPUTHandler(writer http.ResponseWriter, req *http.Request) {

}

func libraryDELETEHandler(writer http.ResponseWriter, req *http.Request) {

}
