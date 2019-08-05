package services

import (
	"fmt"
	"imooc/imooc-iris/repositories"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManager struct {
	repo repositories.MovieRepository
}

func NewMovieServiceManager(repo repositories.MovieRepository) MovieService {
	return &MovieServiceManager{repo: repo}
}

func (m *MovieServiceManager) ShowMovieName() string {
	fmt.Println("我们获取到的视频名称为: " + m.repo.GetMovieName())
	return "我们获取到的视频名称为: " + m.repo.GetMovieName()
}
