package db

import "songs-treasure/pkg/db/model"

type DB interface {
	GetGroup(id string) (group *model.Group, err error)
	GetGroups(group string, page, pageSize uint) (groups []*model.Group, currentPage, pages int64, err error)

	AddSong(group, song, text, releaseDate, link string) (songInfo *SongInfo, err error)
	GetSong(id string) (songInfo *SongInfo, err error)
	GetSongs(group, song, from, to, link string, page, pageSize uint) (songs []*SongInfo, currentPage, pages int64, err error)
	GetSongsByGroupId(id, song, from, to, link string, page, pageSize uint) (songs []*SongInfo, currentPage, pages int64, err error)
	EditSong(id, groupId int, song, release, link string) (songInfo *SongInfo, err error)
	DeleteSong(id int) (err error)

	GetVerses(text string, page, pageSize uint) (songs []*Song, currentPage, pages int64, err error)
	GetVersesBySongId(id string, page, pageSize uint) (songWithVerses *SongWithVerses, currentPage, pages int64, err error)
	EditVerses(id int, text string, versePosition uint) (changedSong *Song, err error)
}
