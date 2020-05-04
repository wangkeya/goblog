package service

import (
	. "goblog/utils/database/mysql"
	"goblog/model"
)

type UserService struct {}

func (u *UserService) CheckLogin(id int) (bool, *model.User, error) {
	var user model.User
	err := DB.First(&user, "id = ?", id).Error
	return  true, &user, err
}


func (u *UserService) GetInfo(id int) ( *model.User, error) {
	var user model.User
	err := DB.First(&user, "id = ?", id).Error
	return  &user, err
}