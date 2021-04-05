package strategy

import (
	"fmt"
	"testing"
)

func TestPostgreSqlDataBase_GenerateStruct(t *testing.T) {
	host := "host=localhost user=postgres password=123456 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db := PostgreSqlDataBase{
		Host: host,
		DB:   nil,
	}
	if err := db.Init(); err == nil {
		tableNames, _ := db.GetTableName("public")
		for _, v := range tableNames {
			fmt.Printf("%s", v)
		}
	}

}
