package service

import (
	"errors"
	"fmt"
	"math/rand"
	"model"
	"time"
	"util"
)

type UserService struct {
}

func (userService *UserService) Register(mobile, plainPassword, nickName, avatar, sex string) (user model.User, err error) {
	var tempUser model.User
	_, err = DbEngine.Where("mobile = ?", mobile).Get(&tempUser)
	if err != nil {
		return tempUser, err
	}

	if tempUser.Id > 0 {
		return tempUser, errors.New("error this mobile has been registered")
	}

	tempUser.Mobile = mobile
	tempUser.Avatar = avatar
	tempUser.Nickname = nickName
	tempUser.Sex = sex
	tempUser.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	tempUser.Passwd = util.MakePwd(plainPassword, tempUser.Salt)
	tempUser.Createat = time.Now()
	tempUser.Token = fmt.Sprintf("%08d", rand.Int31())

	_, err = DbEngine.InsertOne(&tempUser)

	return tempUser, err
}

func (userService *UserService) Login(mobile, plainPassword string) (user model.User, err error) {

	var tempUser model.User
	DbEngine.Where("mobile = ?", mobile).Get(&tempUser)
	if tempUser.Id == 0 {
		return tempUser, errors.New("This user is not exits!")
	}

	if !util.ValidatePassword(plainPassword, tempUser.Salt, tempUser.Passwd) {
		return tempUser, errors.New("Error password!")
	}

	str := fmt.Sprintf("%d", time.Now().Unix())
	token := util.Md5Encode(str)
	tempUser.Token = token
	DbEngine.Id(tempUser.Id).Cols("token").Update(tempUser)
	return tempUser, nil
}
