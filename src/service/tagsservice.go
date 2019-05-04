package service

import (
	"libs/database"
	"model"
)

type TagService struct{}

func (tagService *TagService) GetTagById(tagId int) (*model.Tag, error) {
	db := database.NewMysql()
	defer db.Close()

	var tag model.Tag
	err := db.First(&tag, "id = ?", tagId).Error

	return &tag, err
}
