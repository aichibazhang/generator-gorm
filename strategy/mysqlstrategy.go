package strategy

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MysqlSqlDataBase struct {
	Host string
	DB   *gorm.DB
}

func (pg *MysqlSqlDataBase) Init() (err error) {
	db, _ := gorm.Open(mysql.Open(pg.Host), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	pg.DB = db
	GL_DB = db
	return
}

func (pg *MysqlSqlDataBase) GetTableName(dbname string) (tableNames []string, err error) {
	err = pg.DB.Raw("select table_name as table_name from information_schema.tables where table_schema = ?", dbname).Scan(&tableNames).Error
	return tableNames, err
}

func (pg *MysqlSqlDataBase) GetColumn(tableName string, dbName string) (Columes []Colume, err error) {
	err = pg.DB.Raw("SELECT COLUMN_NAME colume_name,DATA_TYPE data_type,CASE DATA_TYPE WHEN 'longtext' THEN c.CHARACTER_MAXIMUM_LENGTH WHEN 'varchar' THEN c.CHARACTER_MAXIMUM_LENGTH WHEN 'double' THEN CONCAT_WS( ',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE ) WHEN 'decimal' THEN CONCAT_WS( ',', c.NUMERIC_PRECISION, c.NUMERIC_SCALE ) WHEN 'int' THEN c.NUMERIC_PRECISION WHEN 'bigint' THEN c.NUMERIC_PRECISION ELSE '' END AS data_type_long,COLUMN_COMMENT colume_comment FROM INFORMATION_SCHEMA.COLUMNS c WHERE table_name = ? AND table_schema = ?", tableName, dbName).Scan(&Columes).Error
	return Columes, err
}
