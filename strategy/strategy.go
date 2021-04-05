package strategy

import (
	"gorm.io/gorm"
	"html/template"
	"log"
	"os"
	"path"
	"strings"
)


type DataBaseStrategy interface {
	Init() error
	GetTableName(dbname string) (tableNames []string, err error)
	GetColumn(tableName string, dbName string) (Columes []Colume, err error)
}

type DataBase struct {
	strategy DataBaseStrategy
}

// 初始版本自动化代码工具
type AutoCodeStruct struct {
	PackageName string  `json:"packageName"`
	StructName  string  `json:"structName"`
	TableName   string  `json:"tableName"`
	Fields      []Field `json:"fields"`
}

type Field struct {
	FieldName    string `json:"fieldName"`
	FieldType    string `json:"fieldType"`
	DataType     string `json:"dataType"`
	DataTypeLong string `json:"dataTypeLong"`
	Comment      string `json:"comment"`
	ColumnName   string `json:"columnName"`
}

// 第一个参数为模板文件路径,第二个路径为生成模板文件目标地址
func GeneratorStruct(tplPath string, destPath string, dbName string, strategy DataBaseStrategy) {
	tplName := tplPath[(strings.LastIndex(tplPath, "/") + 1):]
	var fileName string
	if strings.Contains(destPath, "/") {
		lastIndex := strings.LastIndex(destPath, "/")
		dirName := destPath[0:lastIndex]
		fileName = destPath[lastIndex+1:]
		if err := os.MkdirAll(dirName, os.ModePerm); err == nil {
			if _, err = os.Stat(dirName + fileName); err == nil {
				if _, err = os.Create(dirName + fileName); err != nil {
					log.Fatal(err)
				}
			}
		}
	} else {
		fileName = destPath
	}
	if tableNames, err := strategy.GetTableName(dbName); err == nil {
		for k, v := range tableNames {
			colums, _ := strategy.GetColumn(v, dbName)
			var fields []Field
			for _, colum := range colums {
				goDataType, _ := GetSysDictionary(colum.DataType)
				field := Field{
					FieldName:    ToHump(colum.ColumeName),
					FieldType:    goDataType.Type,
					DataType:     colum.DataType,
					DataTypeLong: colum.DataTypeLong,
					Comment:      colum.ColumeComment,
					ColumnName:   colum.ColumeName,
				}
				fields = append(fields, field)
			}
			filenameall := path.Base(fileName)
			filesuffix := path.Ext(fileName)
			fileprefix := filenameall[0 : len(filenameall)-len(filesuffix)]
			codeStruct := AutoCodeStruct{
				PackageName: fileprefix,
				StructName:  ToHump(v),
				TableName:   v,
				Fields:      fields,
			}
			// 这里new的参数对应文件名称,parse对应路径
			t, err := template.New(tplName).Funcs(template.FuncMap{
				"haveHeader": func() bool {
					if k < 1 {
						return true
					}
					return false
				},
			}).ParseFiles(tplPath)
			if err != nil {
				log.Fatal("Parse string literal template error:", err)
			}
			file, err := os.OpenFile(destPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
			if err != nil {
				log.Fatal("open failed err:", err)
				return
			}
			err = t.Execute(file, codeStruct)
			if err != nil {
				log.Fatal("Parse string literal template error:", err)
			}
		}
	}
}

type Colume struct {
	ColumeName    string `json:"columeName";gorm:"column:colume_name"`
	DataType      string `json:"dataType";gorm:"column:data_type"`
	DataTypeLong  string `json:"dataTypeLong";gorm:"column:data_type_long"`
	ColumeComment string `json:"columeComment";gorm:"column:colume_comment"`
}

// 将下划线转换为驼峰
func ToHump(filedName string) (filed string) {
	for _, value := range strings.Split(filedName, "_") {
		filed += strings.ToUpper(string(value[0])) + value[1:]
	}
	filed = strings.ToUpper(string(filed[0])) + filed[1:]
	return filed
}

// 下面这些是数据库类型对应go基本数据类型的映射
type SysDictionary struct {
	gorm.Model
	Name                 string                `json:"name" form:"name" gorm:"column:name;comment:'字典名（中）'"`
	Type                 string                `json:"type" form:"type" gorm:"column:type;comment:'字典名（英）'"`
	Status               *bool                 `json:"status" form:"status" gorm:"column:status;comment:'状态'"`
	Desc                 string                `json:"desc" form:"desc" gorm:"column:desc;comment:'描述'"`
	SysDictionaryDetails []SysDictionaryDetail `json:"sysDictionaryDetails" form:"sysDictionaryDetails"`
}

// 如果含有time.Time 请自行import time包
type SysDictionaryDetail struct {
	gorm.Model
	Label           string `json:"label" form:"label" gorm:"column:label;comment:'展示值'"`
	Value           int    `json:"value" form:"value" gorm:"column:value;comment:'字典值'"`
	Status          *bool  `json:"status" form:"status" gorm:"column:status;comment:'启用状态'"`
	Sort            int    `json:"sort" form:"sort" gorm:"column:sort;comment:'排序标记'"`
	SysDictionaryID int    `json:"sysDictionaryID" form:"sysDictionaryID" gorm:"column:sys_dictionary_id;comment:'关联标记'"`
}

// 将数据库类型转换为go类型
func GetSysDictionary(label string) (sysDictionary SysDictionary, err error) {
	var detail SysDictionaryDetail
	GL_DB.Where("label=?", label).First(&detail)
	err = GL_DB.Where("id=?", detail.SysDictionaryID).First(&sysDictionary).Error
	return
}
