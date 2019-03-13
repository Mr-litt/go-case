package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"time"
)

type MysqlConfig struct {
	Host        string
	Port        int
	Username    string
	Password    string
	Database    string
	maxOpen     int
	maxIdle     int
	maxLifetime time.Duration
}

var GMysqlConfig = &MysqlConfig{
	Host:        "127.0.0.1",
	Port:        3306,
	Username:    "root",
	Password:    "12345678",
	Database:    "shortLink",
	maxOpen:     100,
	maxIdle:     10,
	maxLifetime: 60,
}

type Mysql struct {
	db     *sql.DB
	isInit bool
}

var GMysql = &Mysql{}

// 初始化
func (mysql *Mysql) init() error {

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", GMysqlConfig.Username, GMysqlConfig.Password, GMysqlConfig.Host, GMysqlConfig.Port, GMysqlConfig.Database)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return err
	}

	db.SetMaxOpenConns(GMysqlConfig.maxOpen)
	db.SetMaxIdleConns(GMysqlConfig.maxIdle)
	db.SetConnMaxLifetime(GMysqlConfig.maxLifetime * time.Second)
	db.Ping()

	GMysql.db = db
	GMysql.isInit = true
	return nil
}

// 查询
func (mysql *Mysql) Query(querySql string, scan func(*sql.Rows)) error {

	if !mysql.isInit {
		if err := mysql.init(); err != nil {
			return err
		}
	}

	rows, err := mysql.db.Query(querySql)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		scan(rows)
	}

	return nil
}

// 插入
func (mysql *Mysql) Insert(insertSql string, lastInsertId *int64) error {

	if !mysql.isInit {
		if err := mysql.init(); err != nil {
			return err
		}
	}

	res, err := mysql.db.Exec(insertSql)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	*lastInsertId = id

	return nil
}
