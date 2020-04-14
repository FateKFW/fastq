package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/*
DQL：数据库查询语言。关键字：SELECT ... FROM ... WHERE。
DDL ：数据库模式定义语言。关键字：CREATE，DROP，ALTER。
DML：数据操纵语言。关键字：INSERT、UPDATE、DELETE。
DCL：数据控制语言 。关键字：GRANT、REVOKE。
TCL：事务控制语言。关键字：COMMIT、ROLLBACK、SAVEPOINT。
*/

type DBConn struct {
	dbtype string
	ip string
	port string
	name string
	pwd string
	database string
	dql string
}

func main(){
	db := new(DBConn)
	db.initConf()
	if db.dbtype == "mysql" {
		db.opMySQL()
	} else if db.dbtype == "oracle" {
		fmt.Println("Oracle数据库待更新")
	}
}

//初始化连接信息
func (db *DBConn) initConf() *DBConn{
	db.dbtype = "mysql"
	db.ip = "192.168.1.5"
	db.port = "3306"
	db.name = "root"
	db.pwd = "admin"
	db.database = "test"
	db.dql = "select * from test"
	return db
}