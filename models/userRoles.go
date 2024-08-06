package models

// UserRoles [...]
type UserRoles struct {
	UserId string `gorm:"primary_key;column:user_id;type:varchar(255);not null"`
	Users  User   `gorm:"foreignkey:user_id;"`
	RoleId string `gorm:"primary_key;index:IX_UserRoles_RoleId;column:role_id;type:varchar(255);not null"`

	Roles  Role `gorm:"foreignkey:role_id;reference:Id"`
	Record `gorm:"embedded"`
}

// //TableName 数据表名称
// func (m *UserRoles) TableName() string {
// 	return "UserRoles"
// }

func (m *UserRoles) GetID() interface{} {
	return m
}
