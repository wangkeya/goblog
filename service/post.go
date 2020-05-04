package service

import (
	. "goblog/utils/database/mysql"
	"goblog/model"
	"fmt"
	"time"
)

type PostService struct{}

func (postService *PostService) GetPostById(postId int) (*model.Post, error) {
	var post model.Post
	err := DB.First(&post, "id = ?", postId).Error
	return &post, err
}

func (postService *PostService)PageData(page int ,pageRows int) ([]*model.Post,int,error) {
	posts := make([]*model.Post,0)
	var count int
	var offset int
	if page<=1 {
		page = 1
		offset=0
	}else{
		offset = (page-1) * pageRows
	}
	if err := DB.Offset(offset).Limit(pageRows).Order("id desc").Find(&posts).Count(&count).Error; err != nil{
		return posts,count,nil
	}
	return posts,count,nil
}

func (postService *PostService) Insert() error{
	post := new(model.Post)
	post.UserId = 1
	post.Author = "admin"
	post.Title = "测试标题"
	post.Content = "<p>这是测试</p>"
	post.IsTop = 1
	post.UpdateTime = time.Now()

	err := DB.Create(post).Error
	fmt.Println(err)
	return err
}

func (postService *PostService) Update() error {

	post := new(model.Post)
	DB.First(post, "id = ?", 10)

	post.UserId = 1
	post.Author = "admin"
	post.Title = "测试标题"
	post.Content = "<p>这是测试</p>"
	post.IsTop = 1
	post.UpdateTime = time.Now()

	err := DB.Save(post).Error
	return  err
}

func (postService *PostService) Delete(postId int) error {
	post := &model.Post{ID:postId}
	err := DB.Delete(post).Error
	return err
}
