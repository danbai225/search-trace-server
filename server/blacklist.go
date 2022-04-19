package server

import (
	"search-trace-server/db"
	"search-trace-server/model"
)

func BlacklistList(userName string) (res []*model.Blacklist, err error) {
	res = make([]*model.Blacklist, 0)
	err = db.GetDBR().Model(&model.Blacklist{}).Where("username = ?", userName).Scan(&res).Error
	return
}
func BlacklistAdd(blacklist *model.Blacklist) (res *model.Blacklist, err error) {
	tx := db.GetDBW()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Callback()
		}
	}()
	if blacklist.ID == 0 {
		//新建
		err = tx.Create(blacklist).Error
		return blacklist, err
	}
	//修改
	err = tx.Updates(blacklist).Error
	return blacklist, err
}
