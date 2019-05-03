package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pelletier/go-toml"
	"log"
)

func (s *Service) Login(c *gin.Context) (int, interface{}){
	// Load config file
	config, err := toml.LoadFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	email := config.Get("account.email")
	password := config.Get("account.password")

	r := Send("/auth/sign-in", "POST", "", fmt.Sprintf(`{"email":"%s", "password":"%s"}`, email, password))

	return s.makeSuccessJSON(r.Get("data").Get("token").MustString())
}