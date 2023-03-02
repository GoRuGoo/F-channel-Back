package main

import (
	"api/manipulatedb"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	accessPoint = "test_user:pass@tcp(mysql:3306)/kosen?parseTime=true&loc=Asia%2FTokyo"
)

func main() {
	r := gin.Default()
	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{
	// 	"http://localhost:3000",
	// }
	// r.Use(cors.Default())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))
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
		manipulatedb.InsertArticle(db, &postList.Title, &postList.NickName, &postList.KosenName, &postList.Level, &postList.Content)
		c.JSON(http.StatusOK, gin.H{"post": "OK"})
	})
	r.POST("/article/thred/post", func(c *gin.Context) {
		var thredList manipulatedb.Thred
		if err := c.BindJSON(&thredList); err != nil {
			panic(err)
		}
		manipulatedb.InsertThreds(db, &thredList.ThredID, &thredList.NickName, &thredList.Content)
		c.JSON(http.StatusOK, gin.H{"thred": "ok"})
	})
	r.GET("/article/thred/:id", func(c *gin.Context) {
		id := c.Param("id")
		convNum, _ := strconv.Atoi(id)
		c.Data(200, "application/json; charset=utf-8", []byte(manipulatedb.SelectThredsDatabase(db, &convNum)))
	})

	r.Run(":8080")
	defer db.Close()
}
