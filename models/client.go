package models

import "strings"

type Client struct {
	Id             string `json:"id"`
	ClientId       string `json:"client_id" gorm:"unique;column:client_id;type:varchar(200);not null"`
	Secret         string `json:"secret"`
	ClientScope    string `json:"client_scope"`
	ClientName     string `json:"client_name"`
	GrantTypes     string `json:"grant_type"`
	RequireConsent bool   `json:"require_consent"`
	RedirectUris   string `json:"redirect_uris"`
	PostLogoutUris string `json:"post_logout_uris"`
	RequirePKCE    bool   `json:"require_pkce"`
	// 创建时间
	CreatedAt string `json:"created_at"`
	// 更新时间
	UpdatedAt string `json:"updated_at"`
}

func (c *Client) TableName() string {
	return "clients"
}
func (c *Client) GetScopes() []string {
	return strings.Split(c.ClientScope, " ")
}
func (c *Client) GetClientId() string {
	return c.ClientId
}
func (c *Client) GetClientScope() string {
	return c.ClientScope
}
func (c *Client) GetClientName() string {
	return c.ClientName
}
func (c *Client) GetGrantTypes() []string {
	return strings.Split(c.GrantTypes, " ")
}
func (c *Client) GetRequireConsent() bool {
	return c.RequireConsent
}
func (c *Client) GetRedirectUris() []string {
	return strings.Split(c.RedirectUris, " ")
}
func (c *Client) GetPostLogoutUris() []string {
	return strings.Split(c.PostLogoutUris, " ")
}
func (c *Client) GetSecret() string {
	return c.Secret
}
func (c *Client) GetRequirePKCE() bool {
	return c.RequirePKCE
}
