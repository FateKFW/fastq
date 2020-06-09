package main

import (
	"bufio"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
	"io"
	"os"
	"strings"
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
	file string
}

const PROJECT_NAME = "fastq"

func main(){
	//日志模块初始化
	logger.initFastlog(2)

	//执行信息初始化
	db := new(DBConn)
	db.initConf()

	//数据库操作开始
	if db.dbtype == "mysql" {
		db.opMySQL()
	} else if db.dbtype == "sqlite3" {
		db.opSQLite()
	} else {
		logger.error("尚未支持数据库类型")
	}
}

//初始化连接信息
func (db *DBConn) initConf() *DBConn{
	f, _ := os.Open(os.Args[1])
	//f, _ := os.Open("D:/WorkSpace4Idea/fastq/demo/sqlite.txt")
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')

		str := strings.TrimSpace(line)
		key := str[:strings.Index(str,":")]
		value := str[strings.Index(str,":")+1:]

		switch key {
			case "dbtype": db.dbtype = value
			case "ip": db.ip = value
			case "port": db.port = value
			case "name": db.name = value
			case "pwd": db.pwd = value
			case "database": db.database = value
			case "dql": db.dql = value
			case "file": db.file = value
		}

		if err != nil {
			if err == io.EOF {
				break
			} else {
				logger.error(err)
			}
		}
	}
	return db
}

func (db *DBConn) showResult(rows *sql.Rows){
	//读出查询出的列字段名
	cols, _ := rows.Columns()

	//values是每个列的值，这里获取到byte里
	values := make([][]byte, len(cols))

	//query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
	scans := make([]interface{}, len(cols))

	//让每一行数据都填充到[][]byte里面
	for i := range values {
		scans[i] = &values[i]
	}

	var builder strings.Builder

	for _, obj := range cols {
		builder.WriteString(obj+"\t")
	}
	builder.WriteString("\n")

	for rows.Next() {
		if err := rows.Scan(scans...); err != nil {
			panic(err)
		}
		for _, v := range values {
			builder.WriteString(string(v)+"\t")
		}
		builder.WriteString("\n")
	}

	logger.result("执行查询结果",builder.String())
}