package main

import (
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
	"strings"
	"time"
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

	time.Sleep(2 * time.Second)
}

//初始化连接信息
func (db *DBConn) initConf() *DBConn{
	f, _ := os.Open(os.Args[1])
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		param := strings.Split(strings.TrimSpace(line),":")
		switch param[0] {
			case "dbtype": db.dbtype = param[1]
			case "ip": db.ip = param[1]
			case "port": db.port = param[1]
			case "name": db.name = param[1]
			case "pwd": db.pwd = param[1]
			case "database": db.database = param[1]
			case "dql": db.dql = param[1]
		}

		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
	}
	return db
}