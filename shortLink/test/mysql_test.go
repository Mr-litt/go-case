package test

import (
	"testing"
	"log"
	"database/sql"
	"fmt"
	"github.com/mr_litt/go-case/shortLink/app"
)

func TestMysql(t *testing.T) {

	var lastInsertId int64
	insertSql := "insert into `link` (`url`) value ('https://www.baidu.com')"
	err := app.GMysql.Insert(insertSql, &lastInsertId)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Printf("新增id: %d", lastInsertId)

		querySql := fmt.Sprintf("select `url` from `link` where id = %d", lastInsertId)
		err := app.GMysql.Query(querySql, func(rows *sql.Rows) {

			var url string
			rows.Scan( &url)

			log.Printf("url: %s", url)
		})
		if err != nil {
			log.Println(err.Error())
		}
	}
}
