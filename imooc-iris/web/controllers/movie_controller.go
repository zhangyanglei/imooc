package controllers

import (
	"github.com/kataras/iris/mvc"
	"imooc-iris/repositories"
	"imooc-iris/services"
)

type MovieController struct {
}

func (c *MovieController) Get() mvc.View {
	movieRepository := repositories.NewMovieManager()
	movieService := services.NewMovieServiceManager(movieRepository)
	MovieResult := movieService.ShowMovieName()
	return mvc.View{
		Name: "movie/index.html",
		Data: MovieResult,
	}
}
