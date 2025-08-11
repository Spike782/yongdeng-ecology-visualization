package models

import (
	"gorm.io/gorm"
)

type RiskUsage struct {
	ID        int     `gorm:"column:id"`         // 对应表中 id 字段
	Geom      string  `gorm:"column:geom"`       // 对应 geometry 字段（可直接存 WKT 字符串）
	DBId      int64   `gorm:"column:Id"`         // 对应表中 Id 字段（大写）
	Gridcode  int64   `gorm:"column:gridcode"`   // 对应表中 gridcode 字段
	ShapeLeng float64 `gorm:"column:Shape_Leng"` // 对应表中 Shape_Leng 字段
	ShapeArea float64 `gorm:"column:Shape_Area"` // 对应表中 Shape_Area 字段
}

func (RiskUsage) TableName() string {
	return "risk_usages"
}

// 根据 Gridcode 获取地区数据
func GetRiskUsageByGridcode(db *gorm.DB, gridcode int) ([]RiskUsage, error) {
	var riskUsages []RiskUsage
	result := db.Select("*, ST_AsText(geom) as geom").Where("gridcode = ?", gridcode).Find(&riskUsages) // 根据 gridcode 查找
	return riskUsages, result.Error
}
