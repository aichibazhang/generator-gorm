

package base

import "time"

// 自动生成Struct模板

type SpatialRefSys struct{
      Srid  int `gorm:"column:srid;comment:'';type:int4(4)"`
      AuthName  string `gorm:"column:auth_name;comment:'';type:varchar(256)"`
      AuthSrid  int `gorm:"column:auth_srid;comment:'';type:int4(4)"`
      Srtext  string `gorm:"column:srtext;comment:'';type:varchar(2048)"`
      Proj4text  string `gorm:"column:proj4text;comment:'';type:varchar(2048)"`
}

func (SpatialRefSys) TableName() string {
  return "spatial_ref_sys"
}


type PointcloudFormats struct{
      Pcid  int `gorm:"column:pcid;comment:'';type:int4(4)"`
      Srid  int `gorm:"column:srid;comment:'';type:int4(4)"`
      Schema  string `gorm:"column:schema;comment:'';type:text(-5)"`
}

func (PointcloudFormats) TableName() string {
  return "pointcloud_formats"
}


type WindowTest struct{
      Id  int `gorm:"column:id;comment:'';type:int4(4)"`
      Name  string `gorm:"column:name;comment:'';type:text(-5)"`
      Subject  string `gorm:"column:subject;comment:'';type:text(-5)"`
      Score  float64 `gorm:"column:score;comment:'';type:numeric(-5)"`
}

func (WindowTest) TableName() string {
  return "window_test"
}


type Roads struct{
      Gid  int `gorm:"column:gid;comment:'';type:int4(4)"`
      Cat  float64 `gorm:"column:cat;comment:'';type:float8(8)"`
      Label  string `gorm:"column:label;comment:'';type:varchar(80)"`
      Geom  string `gorm:"column:geom;comment:'';type:geometry(16)"`
}

func (Roads) TableName() string {
  return "roads"
}


type ChinaCapital struct{
      Gid  int `gorm:"column:gid;comment:'';type:int4(4)"`
      Name  string `gorm:"column:name;comment:'';type:varchar(100)"`
      Id  float64 `gorm:"column:id;comment:'';type:float8(8)"`
      Level  int `gorm:"column:level;comment:'';type:int2(2)"`
      Geom  string `gorm:"column:geom;comment:'';type:geometry(0)"`
}

func (ChinaCapital) TableName() string {
  return "china_capital"
}


type ChinaCity struct{
      Gid  int `gorm:"column:gid;comment:'';type:int4(4)"`
      Name  string `gorm:"column:name;comment:'';type:varchar(50)"`
      Code  string `gorm:"column:code;comment:'';type:varchar(20)"`
      Pcode  string `gorm:"column:pcode;comment:'';type:varchar(10)"`
      Geom  string `gorm:"column:geom;comment:'';type:geometry(20)"`
}

func (ChinaCity) TableName() string {
  return "china_city"
}


type Company struct{
      Id  int `gorm:"column:id;comment:'';type:int4(4)"`
      Name  string `gorm:"column:name;comment:'公司名称1';type:varchar(20)"`
      Age  int `gorm:"column:age;comment:'';type:int4(4)"`
      Address  string `gorm:"column:address;comment:'';type:varchar(200)"`
      Salary  float64 `gorm:"column:salary;comment:'';type:money(8)"`
}

func (Company) TableName() string {
  return "company"
}


type Company1 struct{
      Id  int `gorm:"column:id;comment:'';type:int4(4)"`
      Name  string `gorm:"column:name;comment:'公司名称';type:varchar(20)"`
      Age  int `gorm:"column:age;comment:'';type:int4(4)"`
      Address  string `gorm:"column:address;comment:'';type:varchar(200)"`
      Salary  float64 `gorm:"column:salary;comment:'';type:money(8)"`
      Bool  *bool `gorm:"column:bool;comment:'';type:bool(1)"`
      Bigint  int `gorm:"column:bigint;comment:'';type:int2(2)"`
      Create  *time.Time `gorm:"column:create;comment:'';type:timestamp(8)"`
      Ct  *time.Time `gorm:"column:ct;comment:'';type:timestamptz(8)"`
      Upd  *time.Time `gorm:"column:upd;comment:'';type:date(4)"`
      Dt  *time.Time `gorm:"column:dt;comment:'';type:time(8)"`
      Ttt  *time.Time `gorm:"column:ttt;comment:'';type:timetz(12)"`
}

func (Company1) TableName() string {
  return "company1"
}


type ChinaCityPoint struct{
      Gid  int `gorm:"column:gid;comment:'';type:int4(4)"`
      Name  string `gorm:"column:name;comment:'';type:varchar(32)"`
      Level  int `gorm:"column:level;comment:'';type:int2(2)"`
      Geom  string `gorm:"column:geom;comment:'';type:geometry(0)"`
}

func (ChinaCityPoint) TableName() string {
  return "china_city_point"
}


type ChinaCounty struct{
      Gid  int `gorm:"column:gid;comment:'';type:int4(4)"`
      Code  string `gorm:"column:code;comment:'';type:varchar(6)"`
      Areaname  string `gorm:"column:areaname;comment:'';type:varchar(100)"`
      Name  string `gorm:"column:name;comment:'';type:varchar(50)"`
      Geom  string `gorm:"column:geom;comment:'';type:geometry(20)"`
}

func (ChinaCounty) TableName() string {
  return "china_county"
}


type ChinaProvince struct{
      Gid  int `gorm:"column:gid;comment:'';type:int4(4)"`
      Name  string `gorm:"column:name;comment:'';type:varchar(50)"`
      Code  string `gorm:"column:code;comment:'';type:varchar(10)"`
      Pcode  string `gorm:"column:pcode;comment:'';type:varchar(10)"`
      Level  int `gorm:"column:level;comment:'';type:int2(2)"`
      Geom  string `gorm:"column:geom;comment:'';type:geometry(20)"`
}

func (ChinaProvince) TableName() string {
  return "china_province"
}
