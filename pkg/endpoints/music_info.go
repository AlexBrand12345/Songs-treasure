package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
	"songs-treasure/internal/config"
	"songs-treasure/pkg/logging"
)

type MusicInfoResponse struct {
	ReleaseDate string
	Text        string
	Link        string
}

func GetMusicsInfo(group, song string) (releaseDate, text, link string, err error) {
	var url string
	var responseStruct MusicInfoResponse

	return "22.10.2020", "blabla\nstrange\n\ntest\n\nanotherone\nend", "", nil

	url = config.MUSIC_INFO_URL + "/info" + fmt.Sprintf("?group=%s&song=%s", group, song)

	resp, err := http.Get(url)
	if err != nil {
		logging.Default.Error(err.Error())
		return
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("Couldn`t connect to music-info server: %v", resp.StatusCode)
		logging.Default.Warnln(err.Error())
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&responseStruct)
	if err != nil {
		logging.Default.Error(err.Error())
		return
	}

	releaseDate = responseStruct.ReleaseDate
	text = responseStruct.Text
	link = responseStruct.Link

	return
}
