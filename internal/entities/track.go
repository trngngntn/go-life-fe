package entities

import "net/http"

type Track struct {
	ID       int
	TrackNum int
	Title    string
	Length   int
}

func GetTrackHandler(writer http.ResponseWriter, req *http.Request) {

}
