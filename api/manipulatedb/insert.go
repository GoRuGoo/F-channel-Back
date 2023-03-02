package manipulatedb

import (
	"database/sql"
	"log"
)

func InsertArticle(db *sql.DB, title, nickName, kosenName *string, level *int, content *string) int64 {
	res, err := db.Exec(
		"INSERT INTO article (title, nick_name,kosen_name,level,content) VALUES (?, ?, ?,?,?)",
		*title, *nickName, *kosenName, *level, *content,
	)
	if err != nil {

		log.Fatalf("insertUser db.Exec error err:%v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatalf("insertUser res.LastInsertId error err:%v", err)
	}
	return id
}

type Thred struct {
	ThredID  int    `json:"thredid"`
	NickName string `json:"nickname"`
	Content  string `json:"content"`
}

func InsertThreds(db *sql.DB, thred_id *int, nickName, content *string) int64 {
	res, err := db.Exec(
		"INSERT INTO threds (thred_id,nick_name,content) VALUES (?,?,?)",
		*thred_id, *nickName, *content,
	)
	if err != nil {
		log.Fatalf("insertThreds db.Exec error err:%v", err)
	}
	id, err := res.LastInsertId()

	if err != nil {
		log.Fatalf("insertThreds res.LastInsertId error err:%v", err)
	}
	return id

}
