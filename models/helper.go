package models

import (
	"github.com/labstack/gommon/log"
	"github.com/wingfeng/idx-oauth2/model"
	"gorm.io/gorm"
)

// Sync2Db 将struct同步数据结构到数据库
func Sync2Db(x *gorm.DB) {
	x.DisableForeignKeyConstraintWhenMigrating = true

	err := x.AutoMigrate(

		new(PersistedGrants),
		new(Role),
		new(User),
		new(UserRoles),
		new(model.Authorization),
		new(Client),
	)
	if err != nil {
		log.Errorf("同步数据结构错误,Error:%v", err)
	}
}
