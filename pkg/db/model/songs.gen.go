// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSong = "songs"

// Song mapped from table <songs>
type Song struct {
	ID          int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	GroupID     int32     `gorm:"column:group_id;not null" json:"group_id"`
	SongName    string    `gorm:"column:song_name;not null" json:"song_name"`
	ReleaseDate time.Time `gorm:"column:release_date;not null" json:"release_date"`
	Link        string    `gorm:"column:link" json:"link"`
}

// TableName Song's table name
func (*Song) TableName() string {
	return TableNameSong
}
