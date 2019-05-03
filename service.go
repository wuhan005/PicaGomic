package main

import "github.com/gin-gonic/gin"

type Service struct {
	Router *gin.Engine
}

func (s *Service) init() {
	s.initRouter()
}