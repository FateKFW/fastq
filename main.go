package main

import (
	"bufio"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
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

type FastQ struct {
	dbtype string
	url string
	port string
	name string
	pwd string
	database string
	dql string
	dml string
}

const PROJECT_NAME = "fastq"

func main(){
	//日志模块初始化
	logger.initFastlog(2)

	//执行信息初始化
	fq := new(FastQ)
	fq.initConf()
	fq.handleOP()
}

//初始化连接信息
func (fq *FastQ) initConf() *FastQ{
	f, _ := os.Open(os.Args[1])
	//f, _ := os.Open("D:/WorkSpace4Idea/fastq/demo/pg.txt")
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')

		str := strings.TrimSpace(line)
		key := str[:strings.Index(str,":")]
		value := str[strings.Index(str,":")+1:]

		switch key {
			case "dbtype": fq.dbtype = value
			case "url": fq.url = value
			case "port": fq.port = value
			case "name": fq.name = value
			case "pwd": fq.pwd = value
			case "database": fq.database = value
			case "dql": fq.dql = value
			case "dml": fq.dml = value
		}

		if err != nil {
			if err == io.EOF {
				break
			} else {
				logger.error(err)
			}
		}
	}
	return fq
}

func (fq *FastQ) handleOP() {
	var db *sql.DB
	var err error

	//获取数据库操作
	if fq.dbtype == "mysql" {			//mysql
		db,err = fq.getMySqlDB()
	} else if fq.dbtype == "sqlite3" {	//sqlite
		db,err = fq.getSQLiteDB()
	} else if fq.dbtype == "postgres" {	//pg
		db,err = fq.getPostgreDB()
	} else {
		logger.error("尚未支持数据库类型")
	}
	defer db.Close()

	if err != nil {
		logger.error(err)
	}

	//执行dql
	if fq.dql != "" {
		fq.handleDQL(db)
	}

	//执行dml
	if fq.dml != "" {
		tx,err := db.Begin()
		if err != nil {
			logger.error(err)
		}
		err = fq.handleDML(db)
		if err != nil {
			tx.Rollback()
			logger.result("DML exec", strings.ReplaceAll(fq.dml, ";", "\n")+"exec failed")
		} else {
			logger.result("DML exec", strings.ReplaceAll(fq.dml, ";", "\n")+"exec success")
			tx.Commit()
		}
	}
}

//执行DQL
func (fq *FastQ) handleDQL(db *sql.DB) {
	ch := make(chan map[string]interface{}, 1000)
	dqlArr := strings.Split(fq.dql, ";")

	//执行查询
	for _,obj := range dqlArr {
		go (func(dql string) {
			rows, err := db.Query(dql)
			if err != nil {
				logger.nerror(err)
			} else {
				res := make(map[string]interface{})
				res["dql"] = dql
				res["rows"] = rows
				ch <- res
			}
		})(obj)
	}

	//结果处理
	for i:=0; i<len(dqlArr); i++ {
		res := <- ch
		fq.showResult(res["dql"].(string), res["rows"].(*sql.Rows))
	}
}

//处理DQL结果
func (fq *FastQ) showResult(dql string, rows *sql.Rows){
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
	if len(cols) > 0 {
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
	}
	logger.result(dql, builder.String())
}

//处理DML
func (fq *FastQ) handleDML(db *sql.DB) error{
	ch := make(chan bool, 1000)
	dmlArr := strings.Split(fq.dml, ";")

	for _,obj := range dmlArr {
		go (func(dml string){
			_,err := db.Exec(dml)
			if err != nil {
				logger.nerror(err)
				ch <- false
			} else {
				ch <- true
			}
		})(obj)
	}

	for i:=0; i<len(dmlArr); i++ {
		if !<-ch {
			return errors.New("")
		}
	}

	return nil
}