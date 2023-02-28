package main

import (
	"api/manipulatedb"
	"database/sql"
	"fmt"
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
		var postList manipulatedb.Article
		if err := c.BindJSON(&postList); err != nil {
			panic(err)
		}
		fmt.Println(postList.Content)
		manipulatedb.InsertArticle(db, &postList.Title, &postList.NickName, &postList.KosenName, &postList.Level, &postList.Content)
	})

	r.Run(":8080")

	defer db.Close()
}
