package repo

import (
	"github.com/wingfeng/idx-gorm/models"
	"github.com/wingfeng/idx/oauth2/model"
	"gorm.io/gorm"
)

type DBClientRepository struct {
	DB *gorm.DB
}

func NewClientRepository(db *gorm.DB) *DBClientRepository {
	return &DBClientRepository{DB: db}
}
func (r *DBClientRepository) GetClientByClientID(id string) (model.IClient, error) {
	result := &models.Client{}
	tx := r.DB.Table(result.TableName()).Where("client_id = ?", id).First(result)
	return result, tx.Error
}
