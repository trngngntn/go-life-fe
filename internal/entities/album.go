package entities

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go-life-fe/internal/components"
	"go-life-fe/internal/utils"
)

type Album struct {
	ID          int    `json:"artist_id"`
	Type        int    `json:"type"`
	Name        string `json:"title"`
	ArtistID    int    `json:"artist_id"`
	ReleaseDate string `json:"release_date"`
	Cover       string
	TrackList   *[]Track
}

func (album *Album) GetTracks() {
	tracks := make([]Track, 17)
	tracks[0] = Track{TrackNum: 1, Title: "Come Together", Length: 259}
	tracks[1] = Track{TrackNum: 2, Title: "Something", Length: 182}
	tracks[2] = Track{TrackNum: 3, Title: "Maxwell's Silver Hammer", Length: 207}
	tracks[3] = Track{TrackNum: 4, Title: "Oh! Darling", Length: 207}
	tracks[4] = Track{TrackNum: 5, Title: "Octopus's Garden", Length: 171}
	tracks[5] = Track{TrackNum: 6, Title: "I Want You (She's So Heavy)", Length: 467}
	tracks[6] = Track{TrackNum: 7, Title: "Here Comes the Sun", Length: 185}
	tracks[7] = Track{TrackNum: 8, Title: "Because", Length: 165}
	tracks[8] = Track{TrackNum: 9, Title: "You Never Give Me Your Money", Length: 243}
	tracks[9] = Track{TrackNum: 10, Title: "Sun King", Length: 146}
	tracks[10] = Track{TrackNum: 11, Title: "Mean Mr. Mustard", Length: 66}
	tracks[11] = Track{TrackNum: 12, Title: "Polythene Pam", Length: 73}
	tracks[12] = Track{TrackNum: 13, Title: "She Came In Through the Bathroom Window", Length: 118}
	tracks[13] = Track{TrackNum: 14, Title: "Golden Slumbers", Length: 91}
	tracks[14] = Track{TrackNum: 15, Title: "Carry That Weight", Length: 96}
	tracks[15] = Track{TrackNum: 16, Title: "The End", Length: 125}
	tracks[16] = Track{TrackNum: 17, Title: "Her Majesty", Length: 23}
	album.TrackList = &tracks
}

func GetAlbumByID(ID int) *Album {
	data := utils.MakeGETRequest(fmt.Sprintf("artists/%d", ID),
		"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxIiwiaWF0IjoxNjQxNTYyMjEzfQ.C3SIrp0ei3KmUXiidd0xGC-P_3lybBgeD9oO15H7yaM")
	var album Album
	json.Unmarshal(data, &album)
	album.GetTracks()
	return &album
}

func GetAlbumHandler(wr http.ResponseWriter, req *http.Request) {
	if ID := utils.GetID(req); ID != -1 {
		components.RenderPage("album", wr, GetAlbumByID(ID), 1)
	}
}
