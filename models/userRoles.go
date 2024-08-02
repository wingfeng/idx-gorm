package models

// UserRoles [...]
type UserRoles struct {
	UserId string `gorm:"primary_key;column:userid;type:varchar(255);not null"`
	Users  User   `gorm:"foreignkey:userid;"`
	RoleId string `gorm:"primary_key;index:IX_UserRoles_RoleId;column:roleid;type:varchar(255);not null"`

	Roles  Role `gorm:"foreignkey:roleid;reference:Id"`
	Record `gorm:"embedded"`
}

// //TableName 数据表名称
// func (m *UserRoles) TableName() string {
// 	return "UserRoles"
// }

func (m *UserRoles) GetID() interface{} {
	return m
}
