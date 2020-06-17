package main

import (
	"database/sql"
	"fmt"
)

func (fq *FastQ) getPostgreDB() (*sql.DB, error){
	db,err := sql.Open(fq.dbtype,
		fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=disable",
			fq.url, fq.port, fq.name, fq.pwd, fq.database))
	if err!=nil {
		return nil,err
	}

	err = db.Ping()
	if err!=nil {
		return nil,err
	}

	//与数据量最大连接数
	db.SetMaxOpenConns(10)
	//最大闲置连接数
	db.SetMaxIdleConns(5)

	return db,nil
}