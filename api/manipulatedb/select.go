package manipulatedb

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"
)

type Article struct {
	PostID    int       `json:"postid"`
	Title     string    `json:"title"`
	NickName  string    `json:"nickname"`
	KosenName string    `json:"kosenname"`
	Level     int       `json:"level"`
	Content   string    `json:"content"`
	Created   time.Time `json:"created"`
	Modified  time.Time `json:"modified"`
}
type Thred struct {
	Id       int       `json:"id"`
	ThredID  int       `json:"thredid"`
	NickName string    `json:"nickname"`
	Content  string    `json:"content"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func SelectDatabase(db *sql.DB) string {
	return (Article{}).getJsonRow(db)
}
func SelectSingleDatabase(db *sql.DB, getRownum *int) string {
	return (Article{}).getJsonSingleRow(db, &getRownum)
}
func SelectThredsDatabase(db *sql.DB) string {
	return (Thred{}).getThredJsonRow(db)
}

func (a Article) getJsonRow(db *sql.DB) string {
	jsonData := []Article{}

	rows, err := db.Query("SELECT * FROM article")

	if err != nil {
		log.Fatalf("connect rows fatal:\n%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&a.PostID, &a.Title, &a.NickName, &a.KosenName, &a.Level, &a.Content, &a.Created, &a.Modified); err != nil {
			log.Fatalf("get rows fatal:\n%v", err)
		}
		jsondata_1 := &Article{PostID: a.PostID, Title: a.Title, NickName: a.NickName, KosenName: a.KosenName, Level: a.Level, Content: a.Content, Created: a.Created, Modified: a.Modified}
		jsonData = append(jsonData, *jsondata_1)
	}
	exportJson, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return string(exportJson)
}

func (a Article) getJsonSingleRow(db *sql.DB, articleId **int) string {
	err := db.QueryRow("SELECT * FROM article WHERE post_id = ?", **articleId).Scan(&a.PostID, &a.Title, &a.NickName, &a.KosenName, &a.Level, &a.Content, &a.Created, &a.Modified)
	if err != nil {
		log.Fatalf("getSingleRow db.QueryRow error:\n%v", err)
	}
	jsonData := &Article{PostID: a.PostID, Title: a.Title, NickName: a.NickName, KosenName: a.KosenName, Level: a.Level, Content: a.Content, Created: a.Created, Modified: a.Modified}
	exportJson, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatal(err)
	}
	return string(exportJson)
}

func (t Thred) getThredJsonRow(db *sql.DB) string {
	jsonData := []Thred{}

	rows, err := db.Query("SELECT * FROM threds ")
	if err != nil {
		log.Fatalf("connect threds rows fatal:\n%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&t.Id, &t.ThredID, &t.NickName, &t.Content, &t.Created, &t.Modified); err != nil {
			log.Fatalf("get threds rosw fatal:\n%v", err)
		}
		jsondat_1 := &Thred{Id: t.Id, ThredID: t.ThredID, NickName: t.NickName, Content: t.Content, Created: t.Created, Modified: t.Modified}
		jsonData = append(jsonData, *jsondat_1)
	}
	exportJson, err := json.Marshal(jsonData)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	return string(exportJson)
}
