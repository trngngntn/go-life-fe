package life

import (
	handlers "go-life-fe/internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var r *mux.Router
var srv *http.Server

func Init() {
	r = mux.NewRouter()
	handlers.InitArtistHandler(r.PathPrefix("/artists").Subrouter())
	handlers.InitLibraryHandler(r.PathPrefix("/library").Subrouter())

	// r.HandleFunc("/login", handlers.LoginGETHandler).Methods("GET")
	// r.HandleFunc("/login", handlers.LoginPOSTHandler).Methods("POST")

	// r.HandleFunc("/artists", entities.GetArtistHandler).Methods("GET")
	// r.HandleFunc("/artists/{ID}", entities.GetArtistHandler).Methods("GET")
	// r.HandleFunc("/albums/{ID}", entities.GetAlbumHandler).Methods("GET")

	r.PathPrefix("/assets/").Handler(
		http.StripPrefix("/assets/", http.FileServer(http.Dir("web/assets"))))

	srv = &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:6969",
	}

}

func Start() {
	log.Fatal(srv.ListenAndServe())
}

func Stop() {
	//srv.Shutdown(nil)
}
