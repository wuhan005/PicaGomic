package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"log"
	"net/url"
	"os"
)

func main(){

	// Load config file
	config, err := toml.LoadFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	email := config.Get("account.email")
	password := config.Get("account.password")

	r := Send("/auth/sign-in", "POST", "", fmt.Sprintf(`{"email":"%s", "password":"%s"}`, email, password))
	token := r.Get("data").Get("token").MustString()
	fmt.Println("Token: " + token)

	categories := Send("/categories", "GET", token, "")
	fmt.Println(categories.Get("data"))

	search := Send("/comics/search?page=1&q=" + url.QueryEscape("亚丝娜"), "GET", token, "")
	fmt.Println(search)

	book := Send("/comics/5b59a1a17ff7c90a607c6374/order/1/pages?page=1", "GET", token, "")
	fmt.Println(book)

	picture := GetImage("/002c342b-6edc-40c2-b8e1-7b4f1efdfd9a.jpg", token)

	dstFile,err := os.Create("pic.png")
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	dstFile.WriteString(picture)
}

