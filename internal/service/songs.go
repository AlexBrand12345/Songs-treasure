package service

import (
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
	Id int `json:""`
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
	for _, song := range songs {
		respSongs = append(respSongs, SongInfo{
			Id:          song.Id,
			GroupId:     song.GroupId,
			Group:       song.Group,
			Song:        song.Song,
			ReleaseDate: song.ReleaseDate,
			Link:        song.Link,
		})
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
	for _, song := range songs {
		respSongs = append(respSongs, SongInfo{
			Id:          song.Id,
			GroupId:     song.GroupId,
			Group:       song.Group,
			Song:        song.Song,
			ReleaseDate: song.ReleaseDate,
			Link:        song.Link,
		})
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

func (srv *service) EditSong(id, group, song, release, link string) (resp EditSongResponse, err error) {
	songInfo, err := srv.db.EditSong(id, group, song, release, link)
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

func (srv *service) DeleteSong(id string) (err error) {
	err = srv.db.DeleteSong(id)
	if err != nil {
		logging.Default.Errorf("Couldn`t delete song, info: %v", err)
		return
	}
	logging.Default.Debugf("Deleted song from DB")

	return
}
