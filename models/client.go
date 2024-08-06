package models

import "strings"

// Client 结构体代表了一个客户端，用于OAuth认证流程。
type Client struct {
	// Id 是客户端的唯一标识符。
	Id string `json:"id"`
	// ClientId 是客户端提供的唯一标识符，用于区分不同的客户端。在数据库中唯一。
	ClientId string `json:"client_id" gorm:"unique;column:client_id;type:varchar(200);not null"`
	// Secret 是客户端的密钥，用于安全验证。
	Secret string `json:"secret"`
	// ClientScope 定义了客户端允许访问的范围。
	ClientScope string `json:"client_scope"`
	// ClientName 是客户端的名称。
	ClientName string `json:"client_name"`
	// GrantTypes 表示客户端支持的授权类型。
	GrantTypes string `json:"grant_type"`
	// RequireConsent 指示用户是否需要对授权进行确认。
	RequireConsent bool `json:"require_consent"`
	// RedirectUris 是客户端用于接收授权码的重定向URI列表。
	RedirectUris string `json:"redirect_uris"`
	// PostLogoutUris 是客户端登出后重定向的URI列表。
	PostLogoutUris string `json:"post_logout_uris"`
	// RequirePKCE 指示是否需要使用PKCE（Proof Key for Code Exchange）进行授权码保护。
	RequirePKCE bool `json:"require_pkce"`
	// CreatedAt 记录了客户端的创建时间。
	CreatedAt string `json:"created_at"`
	// UpdatedAt 记录了客户端的最近更新时间。
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

// GetSecret 方法用于获取当前客户端的密钥。
// 该方法没有参数。
// 返回值是一个字符串切片，包含当前客户端的密钥。
func (c *Client) GetSecret() []string {
	return []string{c.Secret}
}
func (c *Client) GetRequirePKCE() bool {
	return c.RequirePKCE
}
