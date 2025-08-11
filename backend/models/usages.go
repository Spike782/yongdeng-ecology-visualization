package models

import "gorm.io/gorm"

type Usages struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement"`
	Geom      string  `gorm:"column:geom"` // 对应 geometry 字段（可直接存 WKT 字符串）
	Id        int64   `gorm:"column:Id"`
	Gridcode  int64   `gorm:"column:gridcode"`
	ShapeLeng float64 `gorm:"column:Shape_Leng;type:numeric(18,11)"` // 或根据需求用 decimal 类型更精准
	ShapeArea float64 `gorm:"column:Shape_Area;type:numeric(18,11)"`
}

func (Usages) TableName() string {
	return "usages"
}

// 根据 Gridcode 获取地区利用类型数据
func GetUsagesByGridcode(db *gorm.DB, gridcode int) ([]Usages, error) {
	var usages []Usages
	result := db.Select("*, ST_AsText(geom) as geom").Where("gridcode = ?", gridcode).Find(&usages)
	return usages, result.Error
}
