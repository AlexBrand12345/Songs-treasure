package service

import (
	"fmt"
	"songs-treasure/pkg/logging"
)

type Verse struct {
	SongID int    `json:"song_id"`
	Name   string `json:"name"`
	Text   string `json:"text"`
}
type GetVersesResponse struct {
	Search string  `json:"search"`
	Verses []Verse `json:"verses"`
	PaginationResponse
}
type GetVersesByIdResponse struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Verses []Verse `json:"verses"`
	PaginationResponse
}
type EditVersesResponse struct {
	Verse
}

func (srv *service) GetVerses(text string, page, pageSize uint) (resp GetVersesResponse, err error) {
	verses, currentPage, pages, err := srv.db.GetVerses(text, page, pageSize)
	if err != nil {
		logging.Default.Errorf("Couldn`t get verses, info: %v", err)
		return
	}
	logging.Default.Debugf("Got verses from DB")

	respVerses := make([]Verse, len(verses))
	for i, verse := range verses {
		respVerses[i] = Verse{
			SongID: int(verse.SongID),
			Name:   verse.Name,
			Text:   verse.Verses,
		}
	}

	resp = GetVersesResponse{
		Search: text,
		Verses: respVerses,
		PaginationResponse: PaginationResponse{
			PageSize: uint(len(verses)),
			Page:     uint(currentPage),
			Pages:    uint(pages),
		},
	}

	return
}

func (srv *service) GetVersesBySongId(id string, page, pageSize uint) (resp GetVersesByIdResponse, err error) {
	songWithVerses, currentPage, pages, err := srv.db.GetVersesBySongId(id, page, pageSize)
	if err != nil {
		logging.Default.Errorf("Couldn`t get verses, info: %v", err)
		return
	}
	logging.Default.Debugf("Got verses from DB")

	respVerses := make([]Verse, len(songWithVerses.Verses))
	for i, verse := range songWithVerses.Verses {
		respVerses[i] = Verse{
			SongID: int(songWithVerses.SongID),
			Name:   songWithVerses.Name,
			Text:   verse,
		}
	}

	resp = GetVersesByIdResponse{
		Id:     int(songWithVerses.SongID),
		Name:   songWithVerses.Name,
		Verses: respVerses,
		PaginationResponse: PaginationResponse{
			PageSize: uint(len(songWithVerses.Verses)),
			Page:     uint(currentPage),
			Pages:    uint(pages),
		},
	}

	return
}

func (srv *service) EditAllVerses(id int, text string) (resp EditVersesResponse, err error) {
	changedSong, err := srv.db.EditVerses(id, text, 0)
	if err != nil {
		logging.Default.Errorf("Couldn`t edit all verses, info: %v", err)
		return
	}
	logging.Default.Debugf("Got updated song from DB")

	resp = EditVersesResponse{
		Verse: Verse{
			SongID: int(changedSong.SongID),
			Name:   changedSong.Name,
			Text:   changedSong.Verses,
		},
	}

	return
}

func (srv *service) EditVerse(id int, text string, versePosition uint) (resp EditVersesResponse, err error) {
	if versePosition == 0 {
		err = fmt.Errorf("Got wrong verse position - 0, but it starts from 1")
		logging.Default.Error(err)

		return
	}
	changedSong, err := srv.db.EditVerses(id, text, versePosition)
	if err != nil {
		logging.Default.Errorf("Couldn`t edit verse, info: %v", err)
		return
	}
	logging.Default.Debugf("Got updated song from DB")

	resp = EditVersesResponse{
		Verse: Verse{
			SongID: int(changedSong.SongID),
			Name:   changedSong.Name,
			Text:   changedSong.Verses,
		},
	}

	return
}
