
{{if haveHeader}}
package {{.PackageName}}

import (
	"github.com/jinzhu/gorm"
)
// 自动生成Struct模板
{{end}}
type {{.StructName}} struct{
      {{- range .Fields}}
            {{- if eq .FieldType "bool" }}
      {{.FieldName}}  *{{.FieldType}} `gorm:"column:{{.ColumnName}};comment:'{{.Comment}}'{{- if .DataType -}};type:{{.DataType}}{{- if .DataTypeLong -}}({{.DataTypeLong}}){{- end -}}{{- end -}}"`
            {{- else }}
      {{.FieldName}}  {{.FieldType}} `gorm:"column:{{.ColumnName}};comment:'{{.Comment}}'{{- if .DataType -}};type:{{.DataType}}{{- if .DataTypeLong -}}({{.DataTypeLong}}){{- end -}}{{- end -}}"`
            {{- end }}  {{- end }}
}
{{ if .TableName }}
func ({{.StructName}}) TableName() string {
  return "{{.TableName}}"
}
{{ end }}