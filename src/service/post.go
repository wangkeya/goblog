package service

import (
	"libs/database"
	"model"
)

type PostService struct{}

func (postService *PostService) GetPostById(postId int) (*model.Post, error) {
	db := database.NewMysql()
	defer db.Close()

	var post model.Post
	err := db.First(&post, "id = ?", postId).Error

	return &post, err
}
