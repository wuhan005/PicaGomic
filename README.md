# PicaGomic

哔咔哔咔 API - Golang ver.

## 简介

本项目是为了实现哔咔哔咔 PMA 程序而创建的。前端还在开发中…...

如有任何意见或问题，欢迎给本项目提 issue！！



**注意：若无法请求到数据，请确保您已连接上国际互联网。**

## API 参考

### GET `/authorization ` 获取 Token

将`config.sample.toml`另存为`config.toml`，并修改其中的配置为自己的哔咔哔咔账号。



### GET `/categories`  获取本子分类

Header

* `token`：用户的 Token



### GET `/search`  搜索本子

Header

* `token`：用户的 Token

Query

* `keyword`：搜索关键词
* `page`：结果页数(默认为`1`)



### GET `/comic/:id`  获取本子详细信息

Header

- `token`：用户的 Token

Param

* `id`：本子的 Id



### GET  `/episode/:id`  获取本子章节信息

Header

- `token`：用户的 Token

Param

- `id`：本子的 Id



### GET `/content/:bookID/order/:episodeID` 获取本子内容

Header

- `token`：用户的 Token

Param

- `bookID`：本子的 Id
- `episodeID`：章节序号(1, 2, 3...)



### GET `/image/:url`  获取图片(本子内容、封面、分类封面)

Header

- `token`：用户的 Token

Param

- `url`：图片文件名



## 特别感谢

[FlanNep/PicaComic-Api](https://github.com/FlanNep/PicaComic-Api)

