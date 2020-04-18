package main

import (
	"database/sql"
)

func (db *DBConn) getSQLiteConn() (*sql.DB,error){
	return sql.Open(db.dbtype,db.file)
}

func (db *DBConn) opSQLite(){
	conn,err := db.getSQLiteConn()
	if err != nil {
		logger.error(err)
	}

	//执行查询
	rows, err := conn.Query(db.dql)
	if err != nil {
		logger.error(err)
	}

	defer conn.Close()

	db.showResult(rows)
}