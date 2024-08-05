package test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	idxgorm "github.com/wingfeng/idx-gorm"
	"github.com/wingfeng/idx-gorm/models"

	constants "github.com/wingfeng/idx/oauth2/const"
	"github.com/wingfeng/idx/oauth2/utils"
	"gorm.io/gorm"
)

var db *gorm.DB

func initDB() {
	//初始化DB
	//db = utils.GetDB("mysql", "root:kXbXt2nLrL@tcp(localhost:3306)/idx?&parseTime=true")
	db = idxgorm.GetDB("pgx", "host=localhost user=root password=pass@word1 dbname=idx port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	//
	models.Sync2Db(db)
}

func TestSeedData(t *testing.T) {
	//	node, err := snowflake.NewNode(1)
	initDB()

	user := &models.User{}
	user.Id = "7a45cb54-b0ff-4ecd-95b9-074d33aaac1e"
	user.Account = "user1"
	user.DisplayName = "管理员"
	user.Email = "admin@idx.local"

	user.PasswordHash, _ = utils.HashPassword("password1")

	err := db.Save(user).Error
	if err != nil {
		panic(err)
	}
	role := &models.Role{}

	role.Id = "d4d1a7f6-9f33-4ed6-a320-df3754c6e43b"
	role.Name = "SystemAdmin"
	addRole(role)
	addUserRole(user.Id, role.Id)
	role = &models.Role{}

	role.Id = "d4d1a7f6-9f33-4ed6-a320-df3754c6e43c"
	role.Name = "科室主任"
	addRole(role)
	addUserRole(user.Id, role.Id)
	addClient("implicit_client", "secret", "implicit", t)
	addClient("hybrid_client", "secret", "authorization_code implicit", t)
	addClient("code_client", "secret", "authorization_code", t)
	addClient("password_client", "secret", "password", t)
	addClient("local_test", "secret", "authorization_code", t)
	addClient("device_code_client", "secret", string(constants.DeviceCode), t)

}
func TestInsertHybrid(t *testing.T) {
	initDB()
	addClient("oidc-client-implicit.test", "secret", "authorization_code implicit", t)

}
func TestInsertAuthCode(t *testing.T) {
	initDB()
	addClient("local_test", "secret", "authorization_code", t)

}
func TestInsertPasswordClient(t *testing.T) {
	initDB()
	addClient("password_client", "secret", "password", t)

}
func addClient(clientID, secret, grantType string, t *testing.T) {
	//requireSecret := len(secret) > 0
	secretHash, _ := utils.HashPassword(secret)
	client := &models.Client{
		Id:       uuid.NewString(),
		ClientId: clientID,

		ClientName:     fmt.Sprintf("Client %s", clientID),
		Secret:         secretHash,
		GrantTypes:     grantType,
		RedirectUris:   "http://localhost:9000/callback",
		ClientScope:    "openid email profile roles",
		RequireConsent: true,

		//UserSsoLifetime: , can be zero
	}
	//	var count int64
	var result *gorm.DB
	if db.Table("clients").Where("client_id=?", clientID).First(&models.Client{}).RowsAffected > 0 {
		result = db.Table("clients").Updates(client).Where("client_id=?", clientID)

	} else {
		result = db.Table("clients").Save(client).Where("client_id=?", clientID)
	}

	if result.Error != nil {
		t.Logf("insert client error: %v", result.Error)
		panic(result.Error)
	}

}

func addUserRole(uid, rid string) {

	ur := &models.UserRoles{
		RoleId: rid,
		UserId: uid,
	}
	//联合主键的直接用engine来处理
	err := db.Save(ur).Error
	if err != nil {
		panic(err)
	}
}
func addRole(role *models.Role) {

	err := db.Save(role).Error
	if err != nil {
		panic(err)
	}
}
