package models

import "gorm.io/gorm"

type Risks struct {
	ID        int     `gorm:"column:id;primaryKey;autoIncrement"`
	Geom      string  `gorm:"column:geom"` // 对应 geometry 字段（可直接存 WKT 字符串）
	Id        int64   `gorm:"column:Id"`
	Gridcode  int64   `gorm:"column:gridcode"`
	ShapeLeng float64 `gorm:"column:Shape_Leng;type:numeric(18,11)"` // 或根据需求用 decimal 类型更精准
	ShapeArea float64 `gorm:"column:Shape_Area;type:numeric(18,11)"`
}

// TableName 指定关联的数据库表名
func (Risks) TableName() string {
	return "risks"
}

// 根据 Gridcode 获取地区风险度数据
func GetRisksByGridcode(db *gorm.DB, gridcode int) ([]Risks, error) {
	var risks []Risks
	result := db.Select("*, ST_AsText(geom) as geom").Where("gridcode = ?", gridcode).Find(&risks)
	return risks, result.Error
}
