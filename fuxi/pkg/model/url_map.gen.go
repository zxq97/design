// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameURLMap = "url_map"

// URLMap mapped from table <url_map>
type URLMap struct {
	ID       int64     `gorm:"column:id;primaryKey;comment:gen id" json:"id"` // gen id
	ShortURL string    `gorm:"column:short_url;not null" json:"short_url"`
	RealURL  string    `gorm:"column:real_url;not null" json:"real_url"`
	Status   int32     `gorm:"column:status;not null;comment:0:未分配 1:已分配" json:"status"`                    // 0:未分配 1:已分配
	Ctime    time.Time `gorm:"column:ctime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"ctime"`   // 创建时间
	Mtime    time.Time `gorm:"column:mtime;not null;default:CURRENT_TIMESTAMP;comment:最近修改时间" json:"mtime"` // 最近修改时间
}

// TableName URLMap's table name
func (*URLMap) TableName() string {
	return TableNameURLMap
}