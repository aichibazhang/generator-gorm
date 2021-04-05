package strategy

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgreSqlDataBase struct {
	Host string
	DB   *gorm.DB
}

func (pg *PostgreSqlDataBase) Init() (err error) {
	db, err := gorm.Open(postgres.Open(pg.Host), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	pg.DB = db
	GL_DB = db
	return
}

func (pg *PostgreSqlDataBase) GetTableName(dbname string) (tableNames []string, err error) {
	err = pg.DB.Raw("select tablename from pg_tables where schemaname = ?", dbname).Scan(&tableNames).Error
	return tableNames,err
}

func (pg *PostgreSqlDataBase) GetColumn(tableName string, dbName string) (Columes []Colume, err error) {
	err = pg.DB.Raw("SELECT col_description(a.attrelid,a.attnum) as colume_comment, T.typname AS data_type, a.attname as colume_name, ( CASE WHEN A.attlen > 0 THEN A.attlen ELSE A.atttypmod - 4 END ) AS data_type_long FROM pg_class as c,pg_attribute as a,	pg_type T where c.relname = ? and a.attrelid = c.oid and a.attnum>0 and A.atttypid = T.oid", tableName).Scan(&Columes).Error
	return Columes, err
}
