package server

import (
	"fmt"
	"search-trace-server/db"
	"search-trace-server/model"
	"strings"
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
	err = tx.Save(blacklist).Error
	return blacklist, err
}
func BlacklistDelete(u *model.User, id int64) (res *model.Blacklist, err error) {
	tx := db.GetDBW()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Callback()
		}
	}()
	blacklist := model.Blacklist{}
	if u.Role == 1 {
		err = tx.Model(&blacklist).Where("id=?", id).First(&blacklist).Error
	} else {
		err = tx.Model(&blacklist).Where("id=? and username=?", id, u.Name).First(&blacklist).Error
	}
	if err == nil {
		tx.Delete(&blacklist)
	}
	return &blacklist, err
}
func BlacklistAddDomain(domain, uname string) (err error) {
	tx := db.GetDBW()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Callback()
		}
	}()
	domain = fmt.Sprint("*://", domain, "*")
	blacklist := model.Blacklist{}
	err = tx.Model(&blacklist).Where("name=?", "domain").First(&blacklist).Error
	blacklist.Name = "domain"
	blacklist.Mode = 1
	blacklist.UserName = uname
	blacklist.MatchPattern = 1
	b := true
	blacklist.Enable = &b
	split := strings.Split(blacklist.Rules, "\n")
	flag := false
	for _, s := range split {
		if s == domain {
			flag = true
		}
	}
	if !flag {
		split = append(split, domain)
	}
	blacklist.Rules = strings.Join(split, "\n")
	if err != nil && strings.Contains(err.Error(), "record not found") {
		err = tx.Create(&blacklist).Error
	} else if err == nil {
		err = tx.Save(&blacklist).Error
	}
	return err
}
func BlacklistDelDomain(domain, uname string) (err error) {
	tx := db.GetDBW()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Callback()
		}
	}()
	domain = fmt.Sprint("*://", domain, "*")
	blacklist := model.Blacklist{}
	err = tx.Model(&blacklist).Where("name=?", "domain").First(&blacklist).Error
	blacklist.Name = "domain"
	blacklist.Mode = 1
	blacklist.UserName = uname
	blacklist.MatchPattern = 1
	b := true
	blacklist.Enable = &b
	arr := make([]string, 0)
	split := strings.Split(blacklist.Rules, "\n")
	for _, s := range split {
		if s == domain || s == "" {
			continue
		}
		arr = append(arr, s)
	}
	blacklist.Rules = strings.Join(arr, "\n")
	if err != nil && strings.Contains(err.Error(), "record not found") {
		err = tx.Create(&blacklist).Error
	} else if err == nil {
		err = tx.Save(&blacklist).Error
	}
	return err
}
