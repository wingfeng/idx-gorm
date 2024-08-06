package models

import (
	"strings"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

// User 用户信息
type User struct {
	Id                   string    `json:"id" gorm:"primaryKey;column:id;type:varchar(255);not null"`
	Account              string    `json:"account" gorm:"column:account;type:varchar(256)"`
	DisplayName          string    `json:"displayname" gorm:"column:displayname;type:varchar(256)"`
	NormalizedAccount    string    `json:"-" gorm:"unique;column:normalizedaccount;type:varchar(256)"`
	Email                string    `json:"email" gorm:"column:email;type:varchar(256)"`
	NormalizedEmail      string    `json:"-" gorm:"unique;index:EmailIndex;column:normalizedemail;type:varchar(256)"`
	EmailConfirmed       bool      `json:"emailconfirmed" gorm:"column:emailconfirmed;not null"`
	PasswordHash         string    `json:"-" gorm:"column:passwordhash;type:text"`
	SecurityStamp        string    `json:"-" gorm:"column:securitystamp;type:text"`
	PhoneNumber          string    `json:"phonenumber" gorm:"column:phonenumber;type:text"`
	PhoneNumberConfirmed bool      `json:"phonenumberconfirmed" gorm:"column:phonenumberconfirmed;not null"`
	TwoFactorEnabled     bool      `json:"twofactorenabled" gorm:"column:twofactorenabled;not null"`
	IsTemporaryPassword  bool      `json:"istemporarypassword" gorm:"column:istemporarypassword;not null"`
	LockoutEnd           null.Time `json:"lockoutend" gorm:"column:lockoutend"`
	LockoutEnabled       bool      `json:"lockoutenabled" gorm:"column:lockoutenabled;not null"`
	AccessFailedCount    int       `json:"accessfailedcount" gorm:"column:accessfailedcount;type:int;not null"`
	Roles                []Role    `json:"roles" gorm:"many2many:user_roles;foreignkey:id;association_foreignkey:id;"`
	Record               `gorm:"embedded"`
}

// //TableName 数据表名称
// func (m *User) TableName() string {
// 	return "Users"
// }

func (r *User) BeforeCreate(tx *gorm.DB) error {

	r.NormalizedAccount = strings.ToUpper(r.Account)
	r.NormalizedEmail = strings.ToUpper(r.Email)
	return nil
}
func (r *User) BeforeUpdate(tx *gorm.DB) error {

	r.NormalizedAccount = strings.ToUpper(r.Account)
	r.NormalizedEmail = strings.ToUpper(r.Email)
	return nil
}

// Implement model.IUser interface
func (r *User) GetId() string {
	return r.Id
}
func (r *User) GetUserName() string {
	return r.NormalizedAccount
}
func (r *User) GetEmail() string {
	return r.Email
}
func (r *User) GetPasswordHash() string {
	return r.PasswordHash
}
