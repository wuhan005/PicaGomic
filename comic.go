package main

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) Categories(c *gin.Context) (int, interface{}){
	token := c.GetHeader("token")
	data := Send("/categories", "GET", token, "")

	if(data.Get("code").MustInt() != 200){
		return s.makeErrJSON(data.Get("code").MustInt(), data.Get("code").MustInt(), data.Get("message").MustString())
	}

	categories := data.Get("data")
	return s.makeSuccessJSON(categories)
}

func (s *Service) ComicInfo(c *gin.Context) (int, interface{}){
	token := c.GetHeader("token")
	id := c.Param("id")
	data := Send("/comics/" + id, "GET", token,"")

	if(data.Get("code").MustInt() != 200){
		return s.makeErrJSON(data.Get("code").MustInt(), data.Get("code").MustInt(), data.Get("message").MustString())
	}

	return s.makeSuccessJSON(data.Get("data").Get("comic"))
}