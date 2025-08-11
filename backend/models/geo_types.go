package models

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkt"
)

// MultiPolygon 自定义多边形类型，适配 PostGIS
type MultiPolygon struct {
	*geom.MultiPolygon
}

// Scan 从数据库读取数据（WKT格式）并转换为 geom.MultiPolygon
func (m *MultiPolygon) Scan(value interface{}) error {
	if value == nil {
		m.MultiPolygon = nil
		return nil
	}

	// 数据库存储的可能是 []byte 或 string 类型的 WKT
	var wktStr string
	switch v := value.(type) {
	case []byte:
		wktStr = string(v)
	case string:
		wktStr = v
	default:
		return fmt.Errorf("不支持的类型: %T，无法转换为 MultiPolygon", v)
	}

	// 解析 WKT 为 geom 对象
	geomObj, err := wkt.Unmarshal(wktStr)
	if err != nil {
		return fmt.Errorf("解析 WKT 失败: %v", err)
	}

	// 断言为 MultiPolygon 类型
	mp, ok := geomObj.(*geom.MultiPolygon)
	if !ok {
		return errors.New("WKT 内容不是 MultiPolygon 类型")
	}

	m.MultiPolygon = mp
	return nil
}

// Value 将 geom.MultiPolygon 转换为 WKT 字符串存入数据库
func (m MultiPolygon) Value() (driver.Value, error) {
	if m.MultiPolygon == nil {
		return nil, nil
	}
	// 转换为 WKT 格式
	return wkt.Marshal(m.MultiPolygon)
}

// GORM 钩子：在创建表时指定字段类型为 PostGIS 的 geometry
func (MultiPolygon) GormDBDataType(db *gorm.DB, field *gorm.StructField) string {
	// 根据数据库类型返回对应的字段类型，PostGIS 用 geometry
	return "geometry"
}
