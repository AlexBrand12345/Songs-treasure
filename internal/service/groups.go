package service

import "songs-treasure/pkg/logging"

type GetGroupResponse struct {
	Id    int32  `json:"id"`
	Group string `json:"group"`
}
type GetGroupsResponse struct {
	Groups []GetGroupResponse `json:"groups"`
	PaginationResponse
}

func (src *service) GetGroup(id string) (resp GetGroupResponse, err error) {
	group, err := src.db.GetGroup(id)
	if err != nil {
		logging.Default.Errorf("Couldn`t get group, info: %v", err)
		return
	}
	logging.Default.Debugf("Got group from DB")

	resp = GetGroupResponse{
		Id:    group.ID,
		Group: group.GroupName,
	}

	return
}

func (src *service) GetGroups(group string, page, pageSize uint) (resp GetGroupsResponse, err error) {
	groups, currentPage, pages, err := src.db.GetGroups(group, page, pageSize)
	if err != nil {
		logging.Default.Errorf("Couldn`t get groups, info: %v", err)
		return
	}
	logging.Default.Debugf("Got groups from DB")

	respGroups := make([]GetGroupResponse, len(groups))
	for _, group := range groups {
		respGroups = append(respGroups, GetGroupResponse{
			Id:    group.ID,
			Group: group.GroupName,
		})
	}
	resp = GetGroupsResponse{
		Groups: respGroups,
		PaginationResponse: PaginationResponse{
			PageSize: uint(len(groups)),
			Page:     uint(currentPage),
			Pages:    uint(pages),
		},
	}

	return
}
