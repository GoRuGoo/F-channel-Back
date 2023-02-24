package main

import (
	"api/manipulatedb"
	"database/sql"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	accessPoint = "test_user:pass@tcp(mysql:3306)/kosen?parseTime=true&loc=Asia%2FTokyo"
)

func main() {
	r := gin.Default()
	db, err := sql.Open("mysql", accessPoint)
	if err != nil {
		log.Fatalf("first check error:\n%v", err)
	}
	r.GET("/article", func(c *gin.Context) {
		c.Data(200, "application/json; charset=utf-8", []byte(manipulatedb.SelectDatabase(db)))
	})
	r.GET("/article/:id", func(c *gin.Context) {
		id := c.Param("id")
		convNum, _ := strconv.Atoi(id)
		c.Data(200, "application/json; charset=utf-8", []byte(manipulatedb.SelectSingleDatabase(db, &convNum)))
	})
	r.POST("/article/post", func(c *gin.Context) {
		title := c.PostForm("title")
		nickname := c.PostForm("nick_name")
		kosenname := c.PostForm("kosen_name")
		level := c.PostForm("level")
		content := c.PostForm("content")
		manipulatedb.InsertArticle(db, &title, &nickname, &kosenname, &level, &content)
	})
	r.Run(":8080")

	defer db.Close()
}
