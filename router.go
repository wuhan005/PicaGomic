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
	}

	s.Router = r
	err := s.Router.Run(":2334")

	panic(err)
}