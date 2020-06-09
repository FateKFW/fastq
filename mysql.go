package main

import (
	"database/sql"
)

//获取连接
func (db *DBConn) getMySQLConn() (*sql.DB,error){
	return sql.Open(db.dbtype,db.name+":"+db.pwd+"@("+db.ip+":"+db.port+")/"+db.database+"?charset=utf8")
}

func (db *DBConn) opMySQL()  {
	conn,err := db.getMySQLConn()
	if err != nil {
		logger.error(err)
	}

	//执行查询操作
	rows, err := conn.Query(db.dql)
	if err != nil {
		logger.error(err)
	}

	defer conn.Close()	//如果出错或者上述代码执行完毕，延迟关闭连接


	db.showResult(rows)

	//最后得到的map
	/*list := list.New()
	for rows.Next() { //循环，让游标往下推
		if err := rows.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			panic(err)
		}

		row := make(map[string]string) //每行数据

		for k, v := range values { //每行数据是放在values里面，现在把它挪到row里
			row[cols[k]] = string(v)
		}
		list.PushBack(row)
	}*/
	//查询出来的数组
	/*for i := list.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}*/
}