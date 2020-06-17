# fastq
工具说明：无数据库客户端情况下，快速执行sql  
执行方式：打包后将输入文件拖入可执行文件执行

## 配置执行文件
     dbtype:[数据源类型]
     url:[数据源IP或者路径]
     port:[数据源端口]
     name:[用户名]
     pwd:[密码]
     database:[操作哪个databbase]
     dql:[需要执行的dql,多条以";"分割]
     dml:[需要执行的dml,多条以";"分割]
### 配置参考
#### [mysql](https://github.com/FateKFW/fastq/blob/master/demo/mysql.txt)
#### [sqlite](https://github.com/FateKFW/fastq/blob/master/demo/sqlite.txt)
#### [postgresql](https://github.com/FateKFW/fastq/blob/master/demo/pg.txt)
