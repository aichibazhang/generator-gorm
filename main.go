package main

import "github.com/aichibanzhang/fake-data-sql/strategy"

// 这个库用来存取字典,对应go数据类型
//const dbName=`go-model`
const dbName = `public`
const tplPath = "tpl/model.go.tpl"
const destPath = "base/base.go"

func main() {
	//model.InitQueryMysql(dbName)
	//model.InitInsertMysql(isName)
	//model.GeneratorStruct("tpl/model.go.tpl", "base/base.go")
	//base.GetAliClctData()
	host := "host=localhost user=postgres password=123456 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db := strategy.PostgreSqlDataBase{
		Host: host,
		DB:   nil,
	}
	if err := db.Init(); err == nil {
		strategy.GeneratorStruct(tplPath, destPath, dbName, &db)
	}
	//mysqlHost := fmt.Sprintf("root:123456@/%s?charset=utf8&parseTime=True&loc=Local", dbName)
	//mysqlDb := strategy.MysqlSqlDataBase{
	//	Host: mysqlHost,
	//	DB:   nil,
	//}
	//if err := mysqlDb.Init(); err == nil {
	//	strategy.GeneratorStruct(tplPath, destPath, dbName, &mysqlDb)
	//}
}
