package db

import (
	"math"
	"songs-treasure/pkg/db/model"
	"songs-treasure/pkg/logging"
)

func (db *pg) GetGroup(id string) (group *model.Group, err error) {
	err = db.First(group, "id = ?", id).Error
	if err != nil {
		logging.Default.Warn(err.Error())
		return
	}

	return
}

func (db *pg) GetGroups(group string, page, pageSize uint) (groups []*model.Group, currentPage, pages int64, err error) {
	groupQuery := db.Model(&model.Group{})

	if group != "" {
		groupQuery = groupQuery.Where("group_name LIKE ?", "%"+group+"%")
	}

	var rowsCount int64
	err = groupQuery.Count(&rowsCount).Error
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
		groupQuery = groupQuery.Offset(int(offset)).Limit(int(pageSize))
	} else {
		pages = 1
		currentPage = 1
	}

	err = groupQuery.Find(groups).Error
	if err != nil {
		logging.Default.Warn(err.Error())
		return
	}

	return
}
