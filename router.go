package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Service) initRouter() {

	r := gin.Default()

	// CORS Header
	cf := cors.DefaultConfig()
	cf.AllowAllOrigins = true
	cf.AddAllowMethods("POST", "GET", "DELETE", "PUT", "OPTION", "PATCH")
	cf.AddAllowHeaders("token, Content-type, User-Agent")
	r.Use(cors.New(cf))

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

		//GET /content/:bookID/order/:episodeID
		r.GET("/content/:bookID/order/:episodeID", func(c *gin.Context){
			c.JSON(s.ComicDetail(c))
		})

		//GET /image/:url
		r.GET("/image/:url", func(c *gin.Context){
			data := s.ComicMedia(c)
			c.DataFromReader(200, data.ContentLength,"image/png", data.Body, map[string]string{})
		})
	}

	s.Router = r
	err := s.Router.Run(":2334")

	panic(err)
}