package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"songs-treasure/internal/config"
	"songs-treasure/pkg/logging"
)

type MusicInfoRequest struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}
type MusicInfoResponse struct {
	ReleaseDate string
	Text        string
	Link        string
}

func GetMusicsInfo(group, song string) (releaseDate, text, link string, err error) {
	var url string
	var requestStruct MusicInfoRequest
	var responseStruct MusicInfoResponse

	// return "22.10.2020", "blabla\nstrange\n\ntest\n\nanotherone\nend", "", nil

	url = config.MUSIC_INFO_URL + "/info"
	requestStruct = MusicInfoRequest{
		Group: group,
		Song:  song,
	}

	jsonBody, err := json.Marshal(requestStruct)
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		logging.Default.Error(err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
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
