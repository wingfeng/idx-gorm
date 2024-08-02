package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wingfeng/idx-gorm/repo"
)

func TestDBClientRepository_GetClientByClientID(t *testing.T) {
	initDB()
	clientRepo := repo.NewClientRepository(db)
	client, err := clientRepo.GetClientByClientID("test")
	assert.NotNil(t, err)
	assert.Equal(t, "", client.GetClientId())
}
