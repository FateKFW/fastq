package main

import (
	"database/sql"
)

func (fq *FastQ) getSQLiteDB() (*sql.DB, error){
	return sql.Open(fq.dbtype, fq.url)
}