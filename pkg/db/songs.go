package db

import (
	"fmt"
	"math"
	"songs-treasure/pkg/db/model"
	"songs-treasure/pkg/logging"
	"strconv"
	"time"
)

type SongInfo struct {
	Id          int    `gorm:"column:id;primaryKey"`
	GroupId     int    `gorm:"column:group_id;not null"`
	Group       string `gorm:"column:group_name"`
	Song        string `gorm:"column:song_name"`
	ReleaseDate string `gorm:"column:release_date"`
	Link        string `gorm:"column:link"`
}

func (db *pg) AddSong(group, song string) (songInfo *SongInfo, err error) {
	return
}
func (db *pg) GetSong(id string) (songInfo *SongInfo, err error) {
	err = db.Model(&model.Song{}).
		Select("songs.id, songs.group_id, groups.group_name, songs.song_name, songs.release_date, songs.link").
		Joins("JOIN groups ON songs.group_id = groups.id").
		Where("songs.id = ?", id).
		First(&songInfo).Error
	if err != nil {
		logging.Default.Warn(err.Error())
		return
	}

	return
}
func (db *pg) GetSongs(group, song, from, to, link string, page, pageSize uint) (songs []*SongInfo, currentPage, pages int64, err error) {
	songsQuery := db.Model(&model.Song{}).
		Select("songs.id, songs.group_id, groups.group_name, songs.song_name, songs.release_date, songs.link").
		Joins("JOIN groups ON songs.group_id = groups.id")

	if group != "" {
		searchParam := fmt.Sprintf("groups.group_name ILIKE '%s'", "%"+group+"%")
		songsQuery = songsQuery.Where("groups.group_name ILIKE ?", "%"+group+"%").
			Order(searchParam + " DESC")
	}
	if song != "" {
		searchParam := fmt.Sprintf("songs.song_name ILIKE '%s'", "%"+song+"%")
		songsQuery = songsQuery.Where("songs.song_name ILIKE ?", "%"+song+"%").
			Order(searchParam + " DESC")
	}
	if from != "" {
		songsQuery = songsQuery.Where("songs.release_date >= TO_DATE(?, 'DD.MM.YYYY')", from)
	}
	if to != "" {
		songsQuery = songsQuery.Where("songs.release_date <= TO_DATE(?, 'DD.MM.YYYY')", to)
	}
	if link != "" {
		songsQuery = songsQuery.Where("songs.link ILIKE ?", link)
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
func (db *pg) GetSongsByGroupId(id, song, from, to, link string, page, pageSize uint) (songs []*SongInfo, currentPage, pages int64, err error) {
	songsQuery := db.Model(&model.Song{}).
		Select("songs.id, songs.group_id, groups.group_name, songs.song_name, songs.release_date, songs.link").
		Joins("JOIN groups ON songs.group_id = groups.id").
		Where("songs.group_id = ?", id)

	if song != "" {
		searchParam := fmt.Sprintf("songs.song_name ILIKE '%s'", "%"+song+"%")
		songsQuery = songsQuery.Where("songs.song_name ILIKE ?", "%"+song+"%").
			Order(searchParam + " DESC")
	}
	if from != "" {
		songsQuery = songsQuery.Where("songs.release_date >= TO_DATE(?, 'DD.MM.YYYY')", from)
	}
	if to != "" {
		songsQuery = songsQuery.Where("songs.release_date <= TO_DATE(?, 'DD.MM.YYYY')", to)
	}
	if link != "" {
		songsQuery = songsQuery.Where("songs.link ILIKE ?", link)
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
func (db *pg) EditSong(id, group, song, release, link string) (songInfo *SongInfo, err error) {
	var currentSong *model.Song
	err = db.Model(&model.Song{}).
		Where("id = ?", id).
		First(currentSong).Error
	if err != nil {
		logging.Default.Warnf("Couldn`t find song to edit - %v", err.Error())
		return
	}

	var newGroup *model.Group
	if group != "" {
		err = db.Model(&model.Group{}).
			Where("group_name = ?", group).
			First(newGroup).Error
		if err != nil {
			logging.Default.Warnf("Group you wanted to set wasn`t found - %v - %v", group, err.Error())
			return
		}
	}

	if song == "" {
		song = currentSong.SongName
	}
	var releaseDate time.Time
	if release == "" {
		releaseDate = currentSong.ReleaseDate
	} else {
		releaseDate, err = time.Parse("02.01.2006", release)
		if err != nil {
			logging.Default.Warnf("Couldn`t transform date - %v - %v", release, err.Error())
			return
		}
	}
	if link == "" {
		link = currentSong.Link
	}

	err = db.Model(&model.Song{}).
		Updates(model.Song{
			GroupID:     newGroup.ID,
			SongName:    song,
			ReleaseDate: releaseDate,
			Link:        link,
		}).
		Where("id = ?", id).Error
	if err != nil {
		logging.Default.Warn(err.Error())
		return
	}

	return
}
func (db *pg) DeleteSong(id string) (err error) {
	ID, err := strconv.Atoi(id)
	if err != nil {
		logging.Default.Warnf("Couldn`t use id - %v to delete song", err.Error())
	}
	err = db.Model(&model.Song{}).
		Delete(&model.Song{ID: int32(ID)}).Error
	if err != nil {
		logging.Default.Warn(err.Error())
		return
	}

	return
}
