package db

import (
	"fmt"
	"math"
	"songs-treasure/pkg/db/model"
	"songs-treasure/pkg/logging"
	"strings"
	"time"

	"gorm.io/gorm"
)

type SongInfo struct {
	Id          int    `gorm:"column:id;primaryKey"`
	GroupId     int    `gorm:"column:group_id;not null"`
	Group       string `gorm:"column:group_name"`
	Song        string `gorm:"column:song_name"`
	ReleaseDate string `gorm:"column:release_date"`
	Link        string `gorm:"column:link"`
}

func (db *pg) AddSong(group, song, text, releaseDate, link string) (songInfo *SongInfo, err error) {
	err = db.Transaction(func(tx *gorm.DB) (err error) {
		var currentGroup *model.Group
		err = tx.First(&currentGroup, "group_name = ?", group).Error
		if err != nil {
			currentGroup = &model.Group{
				GroupName: group,
			}
			err = tx.Create(&currentGroup).Error
			if err != nil {
				logging.Default.Error(err.Error())
				return
			}
		} else {
			logging.Default.Warnf("Group (%v) already exists", group)
		}

		var currentSongInfo *model.Song
		err = tx.First(&currentSongInfo,
			"group_id = ? AND song_name = ?",
			currentGroup.ID, song).Error
		if err == nil {
			err = fmt.Errorf("This song was added before")
			logging.Default.Warn(err.Error())
			return
		}
		releaseTime, err := time.Parse("02.01.2006", releaseDate)
		if err != nil {
			logging.Default.Warn(err)
			return
		}
		currentSongInfo = &model.Song{
			GroupID:     currentGroup.ID,
			SongName:    song,
			ReleaseDate: releaseTime,
			Link:        link,
		}
		err = tx.Create(&currentSongInfo).Error
		if err != nil {
			err = fmt.Errorf("Couldn`t add song to list")
			logging.Default.Warn(err.Error())
			return
		}

		err = tx.Exec(fmt.Sprintf("INSERT INTO songs_verses (song_id, verses) VALUES (%v, '%v')",
			currentSongInfo.ID, strings.Trim(text, "\n\n"),
		)).Error
		if err != nil {
			err = fmt.Errorf("Couldn`t add song text")
			logging.Default.Warn(err.Error())
			return
		}

		songInfo = &SongInfo{
			Id:          int(currentSongInfo.ID),
			GroupId:     int(currentSongInfo.GroupID),
			Group:       currentGroup.GroupName,
			Song:        currentSongInfo.SongName,
			ReleaseDate: strings.Split(currentSongInfo.ReleaseDate.String(), " ")[0],
			Link:        currentSongInfo.Link,
		}

		return
	})

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
	err = db.Transaction(func(tx *gorm.DB) (err error) {
		songsQuery := tx.Model(&model.Song{}).
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
	})

	return
}
func (db *pg) GetSongsByGroupId(id, song, from, to, link string, page, pageSize uint) (songs []*SongInfo, currentPage, pages int64, err error) {
	err = db.Transaction(func(tx *gorm.DB) (err error) {
		songsQuery := tx.Model(&model.Song{}).
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
	})

	return
}
func (db *pg) EditSong(id, groupId int, song, release, link string) (songInfo *SongInfo, err error) {
	err = db.Transaction(func(tx *gorm.DB) (err error) {
		var currentSong *SongInfo
		err = tx.Model(&model.Song{}).
			Select("songs.id, songs.group_id, groups.group_name, songs.song_name, songs.release_date, songs.link").
			Joins("JOIN groups ON songs.group_id = groups.id").
			Where("songs.id = ?", id).
			First(&currentSong).Error
		if err != nil {
			logging.Default.Warnf("Couldn`t get song to edit - %v", err.Error())
			return
		}

		var newGroup *model.Group
		var group = currentSong.Group
		if groupId != 0 {
			err = tx.Model(&model.Group{}).
				Where("id = ?", groupId).
				First(&newGroup).Error
			if err != nil {
				err = fmt.Errorf("Group you wanted to set wasn`t found - id:%v", groupId)
				logging.Default.Warnf(err.Error())
				return
			}

			groupId = int(newGroup.ID)
			group = newGroup.GroupName
		} else {
			groupId = currentSong.GroupId
		}

		if song == "" {
			song = currentSong.Song
		}

		var equalSong *model.Song
		err = tx.Model(&model.Song{}).
			Select("songs.group_id, songs.song_name").
			Where("songs.group_id = ? AND songs.song_name = ?", groupId, song).
			First(equalSong).Error
		if err == nil {
			logging.Default.Warn(err.Error())
			return
		}

		var releaseDate time.Time
		if release == "" {
			currentSong.ReleaseDate = strings.Split(currentSong.ReleaseDate, "T")[0]
			releaseDate, err = time.Parse("2006-01-02", currentSong.ReleaseDate)
			if err != nil {
				logging.Default.Warnf("Couldn`t transform date - %v - %v", currentSong.ReleaseDate, err.Error())
				return
			}
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

		UpdatedSong := &model.Song{
			GroupID:     int32(groupId),
			SongName:    song,
			ReleaseDate: releaseDate,
			Link:        link,
		}
		err = tx.Model(&model.Song{}).
			Where("id = ?", id).
			Updates(UpdatedSong).Error
		if err != nil {
			logging.Default.Warn(err.Error())
			return
		}

		songInfo = &SongInfo{
			Id:          id,
			GroupId:     int(UpdatedSong.GroupID),
			Group:       group,
			Song:        UpdatedSong.SongName,
			ReleaseDate: strings.Split(UpdatedSong.ReleaseDate.String(), " ")[0],
			Link:        UpdatedSong.Link,
		}

		return
	})

	return
}
func (db *pg) DeleteSong(id int) (err error) {
	resp := db.Model(&model.Song{}).
		Delete(&model.Song{ID: int32(id)})

	err = resp.Error
	rowsDeleted := resp.RowsAffected

	if err != nil {
		err = fmt.Errorf("Couldn`t delete song with id=%v - %v", id, err.Error())
		logging.Default.Warnf(err.Error())
		return
	} else if rowsDeleted > 1 {
		err = fmt.Errorf("Deleted rows count is incorrect - %v (must be 0 or 1)", rowsDeleted)
		logging.Default.Warnf(err.Error())
	} else {
		logging.Default.Warnf("Song with id=%v was found and deleted", id)
	}

	return
}
