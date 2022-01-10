package entities

type Library struct {
	ID          int    `json:"artist_id"`
	Type        int    `json:"type"`
	Name        string `json:"title"`
	ArtistID    int    `json:"artist_id"`
	ReleaseDate string `json:"release_date"`
	Cover       string
	TrackList   *[]Track
}

func (library *Library) GetLibraryTracks(token string) {

}

func GetLibraryByID(ID int, token string) *Album {
	// data := utils.MakeGETRequest(fmt.Sprintf("artists/%d", ID),
	// 	"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxIiwiaWF0IjoxNjQxNTYyMjEzfQ.C3SIrp0ei3KmUXiidd0xGC-P_3lybBgeD9oO15H7yaM")
	// var album Album
	// json.Unmarshal(data, &album)
	// album.GetTracks()
	// return &album
	return nil
}
