package app

import (
	"fmt"
	"log"
	"database/sql"
)

var shortHost = "www.s.cn:8082"

// 短链接路径转长链接
func shortToLong(short string) string {

	// 64进制转10进制
	var url string
	dec := B64ToDec(short)

	// 根据十进制id从mysql获取url
	querySql := fmt.Sprintf("select `url` from `link` where id = %d", dec)
	err := GMysql.Query(querySql, func(rows *sql.Rows) {
		rows.Scan(&url)
	})
	if err != nil {
		log.Printf("查询失败: %s", err.Error())
		return ""
	}
	return url
}

// 长链接转短链接
func longToShort(long string) string {

	// 插入mysql，获取自增id
	var lastInsertId int64
	insertSql := fmt.Sprintf("insert into `link` (`url`) value ('%s')", long)
	err := GMysql.Insert(insertSql, &lastInsertId)
	if err != nil {
		log.Printf("插入失败: %s", err.Error())
		return ""
	}

	// 十进制转64进制
	b64 := DecToB64(int(lastInsertId))

	// 组合短链接域名返回
	return fmt.Sprintf("%s/%s", shortHost, b64)
}
