package service

import (
	"fmt"
	"songs-treasure/pkg/endpoints"
	"songs-treasure/pkg/logging"
	"strconv"
)

type SongInfo struct {
	Id          int    `json:"id"`
	GroupId     int    `json:"group_id"`
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"release_date"`
	Link        string `json:"link"`
}
type AddSongResponse struct {
	SongInfo
}
type GetSongResponse struct {
	SongInfo
}
type GetSongsResponse struct {
	Songs []SongInfo `json:"songs"`
	PaginationResponse
}
type GetSongsByGroupIdResponse struct {
	GroupId int        `json:"group_id"`
	Songs   []SongInfo `json:"songs"`
	PaginationResponse
}
type EditSongResponse struct {
	SongInfo
}

func (srv *service) AddSong(group, song string) (resp AddSongResponse, err error) {
	releaseDate, text, link, err := endpoints.GetMusicsInfo(group, song)
	if err != nil {
		err = fmt.Errorf("Couldn`t get new song description")
		logging.Default.Error(err.Error())
	}
	songInfo, err := srv.db.AddSong(group, song, text, releaseDate, link)
	if err != nil {
		logging.Default.Errorf("Couldn`t add song, info: %v", err)
		return
	}
	logging.Default.Debugf("Added song to DB")

	resp = AddSongResponse{
		SongInfo: SongInfo{
			Id:          songInfo.Id,
			GroupId:     songInfo.GroupId,
			Group:       songInfo.Group,
			Song:        songInfo.Song,
			ReleaseDate: songInfo.ReleaseDate,
			Link:        songInfo.Link,
		},
	}

	return
}

func (srv *service) GetSong(id string) (resp GetSongResponse, err error) {
	songInfo, err := srv.db.GetSong(id)
	if err != nil {
		logging.Default.Errorf("Couldn`t get song, info: %v", err)
		return
	}
	logging.Default.Debugf("Got song from DB")

	resp = GetSongResponse{
		SongInfo: SongInfo{
			Id:          songInfo.Id,
			GroupId:     songInfo.GroupId,
			Group:       songInfo.Group,
			Song:        songInfo.Song,
			ReleaseDate: songInfo.ReleaseDate,
			Link:        songInfo.Link,
		},
	}

	return
}

func (srv *service) GetSongs(group,
	song, from, to, link string, page, pageSize uint) (resp GetSongsResponse, err error) {
	songs, currentPage, pages, err := srv.db.GetSongs(group, song, from, to, link, page, pageSize)
	if err != nil {
		logging.Default.Errorf("Couldn`t get song, info: %v", err)
		return
	}
	logging.Default.Debugf("Got songs from DB")

	respSongs := make([]SongInfo, len(songs))
	for i, song := range songs {
		respSongs[i] = SongInfo{
			Id:          song.Id,
			GroupId:     song.GroupId,
			Group:       song.Group,
			Song:        song.Song,
			ReleaseDate: song.ReleaseDate,
			Link:        song.Link,
		}
	}
	resp = GetSongsResponse{
		Songs: respSongs,
		PaginationResponse: PaginationResponse{
			PageSize: uint(len(songs)),
			Page:     uint(currentPage),
			Pages:    uint(pages),
		},
	}

	return
}

func (srv *service) GetSongsByGroupId(groupId,
	song, from, to, link string, page, pageSize uint) (resp GetSongsByGroupIdResponse, err error) {
	songs, currentPage, pages, err := srv.db.GetSongsByGroupId(groupId, song, from, to, link, page, pageSize)
	if err != nil {
		logging.Default.Errorf("Couldn`t get song, info: %v", err)
		return
	}
	logging.Default.Debugf("Got songs from DB")

	respSongs := make([]SongInfo, len(songs))
	for i, song := range songs {
		respSongs[i] = SongInfo{
			Id:          song.Id,
			GroupId:     song.GroupId,
			Group:       song.Group,
			Song:        song.Song,
			ReleaseDate: song.ReleaseDate,
			Link:        song.Link,
		}
	}

	groupIdInt, _ := strconv.Atoi(groupId)
	resp = GetSongsByGroupIdResponse{
		GroupId: groupIdInt,
		Songs:   respSongs,
		PaginationResponse: PaginationResponse{
			PageSize: uint(len(songs)),
			Page:     uint(currentPage),
			Pages:    uint(pages),
		},
	}

	return
}

func (srv *service) EditSong(id, groupId int, song, release, link string) (resp EditSongResponse, err error) {
	songInfo, err := srv.db.EditSong(id, groupId, song, release, link)
	if err != nil {
		logging.Default.Errorf("Couldn`t edit song, info: %v", err)
		return
	}
	logging.Default.Debugf("Got updated song from DB")

	resp = EditSongResponse{
		SongInfo: SongInfo{
			Id:          songInfo.Id,
			GroupId:     songInfo.GroupId,
			Group:       songInfo.Group,
			Song:        songInfo.Song,
			ReleaseDate: songInfo.ReleaseDate,
			Link:        songInfo.Link,
		},
	}

	return
}

func (srv *service) DeleteSong(id int) (err error) {
	err = srv.db.DeleteSong(id)
	if err != nil {
		logging.Default.Errorf("Couldn`t delete song, info: %v", err)
		return
	}
	logging.Default.Debugf("Deleted song from DB")

	return
}
