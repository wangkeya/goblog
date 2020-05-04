package service

import (
	. "goblog/utils/database/mysql"
	"goblog/model"
)

type TagService struct{}

func (tagService *TagService) GetTagById(tagId int) (*model.Tag, error) {
	var tag model.Tag
	err := DB.First(&tag, "id = ?", tagId).Error
	return &tag, err
}

