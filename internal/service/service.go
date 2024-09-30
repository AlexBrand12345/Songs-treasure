package service

import (
	"songs-treasure/pkg/db"
)

type Service interface {
	GetGroup(id string) (resp GetGroupResponse, err error) // by id
	GetGroups(group string, page, pageSize uint) (resp GetGroupsResponse, err error)

	AddSong(group, song string) (resp AddSongResponse, err error)
	GetSong(id string) (resp GetSongResponse, err error) // by id
	GetSongs(group, song, from, to, link string, page, pageSize uint) (resp GetSongsResponse, err error)
	GetSongsByGroupId(id, song, from, to, link string, page, pageSize uint) (resp GetSongsByGroupIdResponse, err error) // by group id
	EditSong(id, groupId int, song, release, link string) (resp EditSongResponse, err error)
	DeleteSong(id int) (err error)

	GetVerses(text string, page, pageSize uint) (resp GetVersesResponse, err error)
	GetVersesBySongId(id string, page, pageSize uint) (resp GetVersesByIdResponse, err error)
	EditVerse(id int, text string, versePosition uint) (resp EditVersesResponse, err error)
	EditAllVerses(id int, text string) (resp EditVersesResponse, err error)
}
type service struct {
	db db.DB
}

func NewService(db db.DB) *service {
	return &service{
		db: db,
	}
}
