package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wingfeng/idx-gorm/repo"
)

func TestGetUser(t *testing.T) {
	initDB()
	repo := repo.NewUserRepository(db)
	user, err := repo.GetUserByName("user1")
	assert.Nil(t, err)
	t.Logf("User %v", user)
	//	assert.Equal(t, user.(models.User).Roles[0].Name, "SystemAdmin")

}
