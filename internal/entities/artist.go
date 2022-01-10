package entities

import (
	"encoding/json"
	"fmt"

	"go-life-fe/internal/utils"
)

type Artist struct {
	ID           int    `json:artist_id`
	Name         string `json:"display_name"`
	Albums       []Album
	Singles      []Album
	Compilations []Album
}

func (artist *Artist) GetArtistAlbums(token string) {

}

func GetArtistByID(ID int, token string) *Artist {
	data := utils.MakeGETRequest(fmt.Sprintf("artists/%d", ID), token)
	var artist Artist
	json.Unmarshal(data, &artist)
	artist.GetArtistAlbums(token)
	return &artist
}

func GetArtistList(token string) *[]Artist {
	data := utils.MakeGETRequest("artists", token)
	var artistList []Artist
	json.Unmarshal(data, &artistList)
	return &artistList
}
