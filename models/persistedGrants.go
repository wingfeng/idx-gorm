package models

// PersistedGrants [...]
type PersistedGrants struct {
	Principal string `gorm:"primaryKey;column:principal;type:varchar(200)"`
	ClientId  string `gorm:"primaryKey;column:client_id;type:varchar(200);not null"`

	Scope  string `gorm:"column:scope;type:text;not null"`
	Record `gorm:"embedded"`
}

func (m *PersistedGrants) TableName() string {
	return "persisted_grants"
}
