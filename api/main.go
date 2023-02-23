package main

import (
	"api/manipulatedb"
	"database/sql"
	"log"

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
	r.Run()

	defer db.Close()
}
