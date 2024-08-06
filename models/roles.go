package models

import (
	"strings"

	"gorm.io/gorm"
)

// Role [...]
type Role struct {
	Id             string `json:"-" gorm:"primary_key;column:id;type:varchar(255);not null"`
	Name           string `gorm:"column:name;type:varchar(256)"`
	NormalizedName string `json:"-" gorm:"unique;column:normalizedname;type:varchar(256)"`
	Users          []User `json:"-" gorm:"many2many:user_roles;foreignkey:id;association_foreignkey:id;"`
	Record         `json:"-" gorm:"embedded"`
}

// //TableName 数据表名称
// func (m *Role) TableName() string {
// 	return "Roles"
// }

// SetID 获取当前记录的ID
func (r *Role) SetID(v interface{}) {
	r.Id = v.(string)
}

func (r *Role) GetID() interface{} {
	return r.Id
}
func (r *Role) BeforeCreate(tx *gorm.DB) error {

	r.NormalizedName = strings.ToUpper(r.Name)

	return nil
}
func (r *Role) BeforeUpdate(tx *gorm.DB) error {

	r.NormalizedName = strings.ToUpper(r.Name)

	return nil
}
