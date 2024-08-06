package repo

import (
	"strings"

	"github.com/wingfeng/idx-gorm/models"
	"github.com/wingfeng/idx-oauth2/model"
	"gorm.io/gorm"
)

type DBUserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *DBUserRepository {
	return &DBUserRepository{DB: db}
}
func (ur *DBUserRepository) GetUser(userId string) (model.IUser, error) {
	var user models.User
	ur.DB.SetupJoinTable(&models.User{}, "Roles", &models.UserRoles{})
	tx := ur.DB.Where("id = ?", userId).Preload("Roles").First(&user)
	return &user, tx.Error
}
func (ur *DBUserRepository) GetUserByName(username string) (model.IUser, error) {
	var user models.User
	ur.DB.SetupJoinTable(&models.User{}, "Roles", &models.UserRoles{})
	tx := ur.DB.Where("normalizedaccount = ?", strings.ToUpper(username)).Preload("Roles").First(&user)
	return &user, tx.Error
}
