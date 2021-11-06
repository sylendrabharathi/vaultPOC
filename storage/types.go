package storage

import "github.com/hashicorp/vault/api"

type Vault struct {
	Client *api.Client
}

type V2APIKey struct {
	ApiKey   string `json:"apiKey"`
	Email    string `json:"email"`
	Active   bool   `json:"active"`
	Comments string `json:"comments"`
}

type V3APIKey struct {
	ApiKey      string   `json:"apiKey"`
	Email       string   `json:"email"`
	User        string   `json:"user"`
	Active      bool     `json:"active"`
	Comments    string   `json:"comments"`
	Permissions []string `json:"permissions"`
}

type V4APIKey struct {
	ApiKey      string   `json:"apiKey"`
	Email       string   `json:"email"`
	User        string   `json:"user"`
	Active      bool     `json:"active"`
	CanSendAll  bool     `json:"canSendAll"`
	Comments    string   `json:"comments"`
	Permissions []string `json:"permissions"`
}
