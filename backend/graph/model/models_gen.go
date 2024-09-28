// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/hawa130/serverx/ent"
)

type BatchCGroup struct {
	Data []*CGroup `json:"data,omitempty"`
	Ok   bool      `json:"ok"`
}

type BatchCPolicy struct {
	Data []*CPolicy `json:"data,omitempty"`
	Ok   bool       `json:"ok"`
}

type CGroup struct {
	Sub string `json:"sub"`
	Obj string `json:"obj"`
}

type CGroupInput struct {
	Sub string `json:"sub"`
	Obj string `json:"obj"`
}

type CGroupResult struct {
	Sub string `json:"sub"`
	Obj string `json:"obj"`
	Ok  bool   `json:"ok"`
}

type CPolicy struct {
	Sub string `json:"sub"`
	Obj string `json:"obj"`
	Act string `json:"act"`
}

type CPolicyResult struct {
	Sub string `json:"sub"`
	Obj string `json:"obj"`
	Act string `json:"act"`
	Ok  bool   `json:"ok"`
}

type CRequestInput struct {
	Sub string `json:"sub"`
	Obj string `json:"obj"`
	Act string `json:"act"`
}

type LoginInput struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginPayload struct {
	Token *string   `json:"token,omitempty"`
	User  *ent.User `json:"user,omitempty"`
}

type RegisterInput struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UpdateCGroup struct {
	New *CGroup `json:"new"`
	Old *CGroup `json:"old"`
	Ok  bool    `json:"ok"`
}

type UpdateCPolicy struct {
	New *CPolicy `json:"new"`
	Old *CPolicy `json:"old"`
	Ok  bool     `json:"ok"`
}

type UpdatePasswordInput struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
