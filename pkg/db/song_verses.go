package db

import (
	"fmt"
	"math"
	"regexp"
	"songs-treasure/pkg/db/model"
	"songs-treasure/pkg/logging"
	"strings"
)

// select verses, tsv, rank
// from songs_verses,
//     to_tsquery('simple', 'the|ke') query,
//     ts_rank(tsv, query) rank
// where tsv @@ query
//     or verses ILIKE '%the ke%'
// ORDER BY verses ILIKE '%the ke%' desc, rank desc;

type Song struct {
	SongID uint   `gorm:"primary_key;not null"`
	Name   string `gorm:"column:song_name;not null"`
	Verses string
}
type SongWithVerses struct {
	SongID uint   `gorm:"primary_key;not null"`
	Name   string `gorm:"column:song_name;"`
	Verses []string
}

func (db *pg) GetVerses(text string, page, pageSize uint) (songs []*Song, currentPage, pages int64, err error) {
	songsQuery := db.Model(&model.SongsVerse{}).
		Select("songs_verses.song_id, songs.song_name, songs_verses.verses").
		Joins("JOIN songs ON songs.id = songs_verses.song_id")

	if text != "" {
		re := regexp.MustCompile(`\s+`)
		text = re.ReplaceAllString(text, " ")

		query := fmt.Sprintf("to_tsquery('simple', '%s')", strings.Trim(strings.ReplaceAll(text, " ", "|"), "|"))
		songsQuery = songsQuery.
			Where("songs_verses.verses ILIKE ?", "%"+text+"%").
			Or(fmt.Sprintf("songs_verses.tsv @@ %s", query)).
			Order(fmt.Sprintf("songs_verses.verses ILIKE '%s' DESC, ts_rank(songs_verses.tsv, %s) DESC", "%"+text+"%", query))
	}

	var rowsCount int64
	err = songsQuery.Count(&rowsCount).Error
	if err != nil {
		logging.Default.Warn(err.Error())
		return
	}

	if page > 0 && pageSize > 0 {
		pages = int64(math.Ceil(float64(rowsCount) / float64(pageSize)))
		if int64(page) > pages {
			currentPage = pages
		} else {
			currentPage = int64(page)
		}

		offset := (currentPage - 1) * int64(pageSize)
		songsQuery = songsQuery.Offset(int(offset)).Limit(int(pageSize))
	} else {
		pages = 1
		currentPage = 1
	}

	err = songsQuery.Find(&songs).Error
	if err != nil {
		logging.Default.Warn(err.Error())
		return
	}

	return
}

func (db *pg) GetVersesBySongId(id string, page, pageSize uint) (songWithVerses *SongWithVerses, currentPage, pages int64, err error) {
	var song *Song
	songsQuery := db.Model(&model.SongsVerse{}).
		Select("songs_verses.song_id, songs.song_name, songs_verses.verses").
		Joins("JOIN songs ON songs.id = songs_verses.song_id").
		Where("songs_verses.song_id = ?", id).
		Limit(1)

	err = songsQuery.First(&song).Error
	if err != nil {
		logging.Default.Warn(err.Error())
		return
	}

	songVerses := strings.Split(strings.Trim(song.Verses, "\n\n"), "\n\n")

	if page > 0 && pageSize > 0 {
		pages = int64(math.Ceil(float64(len(songVerses)) / float64(pageSize)))
		if int64(page) > pages {
			currentPage = pages
		} else {
			currentPage = int64(page)
		}

		offset := (currentPage - 1) * int64(pageSize)
		songVerses = songVerses[offset:int64(math.Min(float64(offset+int64(pageSize)), float64(len(songVerses))))]
	} else {
		pages = 1
		currentPage = 1
	}

	songWithVerses = &SongWithVerses{
		SongID: song.SongID,
		Name:   song.Name,
		Verses: songVerses,
	}

	return
}

func (db *pg) EditVerses(id int, text string, versePosition uint) (changedSong *Song, err error) {
	var song *Song
	songsQuery := db.Model(&model.SongsVerse{}).
		Select("songs_verses.song_id, songs.song_name, songs_verses.verses").
		Joins("JOIN songs ON songs.id = songs_verses.song_id").
		Where("songs_verses.song_id = ?", id).
		Limit(1)

	err = songsQuery.First(&song).Error
	if err != nil {
		logging.Default.Warn(err.Error())
		return
	}

	songVerses := strings.Split(strings.Trim(song.Verses, "\n\n"), "\n\n")

	changedText := strings.Split(strings.Trim(text, "\n\n"), "\n\n")
	if changedText[0] == "" {
		err = fmt.Errorf("Song`s verse(s) can`t be deleted")
		logging.Default.Error(err)

		return
	}

	if versePosition == 0 {
		songVerses = changedText
	} else {
		changedVerse := int(math.Min(
			math.Max(float64(1), float64(versePosition)),
			float64(len(songVerses)),
		))

		songVerses[changedVerse-1] = changedText[0]
	}

	changedSong = &Song{
		SongID: song.SongID,
		Name:   song.Name,
		Verses: strings.Join(songVerses, "\n\n"),
	}
	err = db.Model(&model.SongsVerse{}).
		Where("song_id = ?", song.SongID).
		Update("verses", strings.Join(songVerses, "\n\n")).Error
	if err != nil {
		logging.Default.Warn(err.Error())
		return
	}

	return
}
