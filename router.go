package main

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) initRouter() {

	r := gin.Default()

	{
		//GET /authorization
		r.GET("/authorization", func(c *gin.Context) {
			c.JSON(s.Login(c))
		})
	}

	{
		//GET /categories
		r.GET("/categories", func(c *gin.Context){
			c.JSON(s.Categories(c))
		})

		//GET /search
		r.GET("/search", func(c *gin.Context){
			c.JSON(s.SearchKeyword(c))
		})
	}

	{
		//GET /comic/:id
		r.GET("/comic/:id", func(c *gin.Context){
			c.JSON(s.ComicInfo(c))
		})

		//GET /episode/:id
		r.GET("/episode/:id", func(c *gin.Context){
			c.JSON(s.ComicEpisode(c))
		})
	}

	s.Router = r
	err := s.Router.Run(":2334")

	panic(err)
}