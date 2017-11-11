package model

type Role struct {
	User       string `json:"user"`
	Role       string `json:"role"`
	Permission string `json:"permission"`
}

type RolesResponse struct {
	Roles []string `json:"roles"`
}