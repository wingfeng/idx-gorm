package test

import (
	"testing"

	"github.com/wingfeng/idx-gorm/repo"
)

func TestSaveConsents(t *testing.T) {

	initDB()
	repo := &repo.DBConsentRepository{DB: db}

	clientId := "test_client"
	principal := "test_principal"
	scopes := []string{"scope1", "scope2"}

	err := repo.SaveConsents(clientId, principal, scopes)
	if err != nil {
		t.Errorf("SaveConsents() error = %v", err)
	}

}
