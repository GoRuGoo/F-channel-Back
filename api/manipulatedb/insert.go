package manipulatedb

import (
	"database/sql"
	"log"
)

func InsertArticle(db *sql.DB, title, nickName, kosenName *string, level *string, content *string) int64 {
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
