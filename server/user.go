package server

import (
	"errors"
	"fmt"
	"github.com/gogf/guuid"
	"search-trace-server/db"
	"search-trace-server/model"
	"search-trace-server/utils"
	"strings"
	"time"
)

const (
	userTokenCachePrefix = "user_token_"
	tokenCacheTime       = time.Hour * 24 * 7
	userCachePrefix      = "user_"
	userCacheTime        = time.Hour * 24
	encryption           = "search-trace-server"
)

/*
密码加密
*/
func passwordEncryption(pass string) string {
	return utils.Sha256(utils.Sha256(pass) + encryption)
}

// UserGetUserByToken 根据token获取用户信息
func UserGetUserByToken(token string) (user *model.User, err error) {
	key := fmt.Sprint(userTokenCachePrefix, token)
	value, err := db.GetCache().Value(key)
	if err != nil {
		return
	}
	name := value.Data().(string)
	user, err = UserGetUserFromCache(name)
	return
}

// UserGetUserFromCache 从缓存中获取用户信息
func UserGetUserFromCache(name string) (user *model.User, err error) {
	key := fmt.Sprint(userCachePrefix, name)
	value, err := db.GetCache().Value(key)
	if err != nil {
		user, err = UserGetOneByName(name)
		if err == nil {
			db.GetCache().Add(key, userCacheTime, user)
		}
		return
	}
	user = value.Data().(*model.User)
	return
}

// LoginByPass 通过用户名密码登陆
func LoginByPass(name, pass string) (token string, user *model.User, err error) {
	user, err = UserGetOneByName(name)
	if err != nil {
		return
	}
	if passwordEncryption(pass) != user.Pass {
		err = errors.New("验证未通过")
		return
	}
	token = strings.ReplaceAll(guuid.New().String(), "-", "")
	key := fmt.Sprint(userTokenCachePrefix, token)
	db.GetCache().Add(key, tokenCacheTime, user.Name)
	return
}

// UserGetOneByName 从数据库中查询用户根据用户名
func UserGetOneByName(name string) (user *model.User, err error) {
	tx := db.GetDB()
	defer func() {
		if err != nil {
			tx.Callback()
		} else {
			tx.Commit()
		}
	}()
	u := model.User{}
	err = tx.Model(&u).Where("name=?", name).First(&u).Error
	if err != nil {
		return
	}
	user = &u
	return
}

// UserCreate 创建用户
func UserCreate(user *model.User) (err error) {
	tx := db.GetDB()
	defer func() {
		if err != nil {
			tx.Callback()
		} else {
			tx.Commit()
		}
	}()

	user.Pass = passwordEncryption(user.Pass)
	err = tx.Create(user).Error
	return
}
