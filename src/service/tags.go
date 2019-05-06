package service

import (
	. "libs/database/mysql"
	"model"
)

type TagService struct{}

func (tagService *TagService) GetTagById(tagId int) (*model.Tag, error) {
	var tag model.Tag
	err := DB.First(&tag, "id = ?", tagId).Error
	return &tag, err
}

