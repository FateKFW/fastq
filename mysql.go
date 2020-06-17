package main

import (
	"database/sql"
	"fmt"
)

//获取连接
func (fq *FastQ) getMySqlDB() (*sql.DB, error){
	return sql.Open(fq.dbtype, fmt.Sprintf("%s:%s@(%s:%v)/%s?charset=utf8",
		fq.name, fq.pwd, fq.url, fq.port, fq.database))
}