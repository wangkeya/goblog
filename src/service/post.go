package service

import (
	. "libs/database/mysql"
	"model"
)

type PostService struct{}

func (postService *PostService) GetPostById(postId int) (*model.Post, error) {
	defer CloseDb()

	var post model.Post
	err := DB.First(&post, "id = ?", postId).Error

	return &post, err
}
