package entity

import "janus-portal/module/command/role_permission/domain/valobj"

type PermissionEntity struct {
	code        *valobj.Code
	name        string
	description string
}

func NewPermissionEntity(
	code *valobj.Code,
	name string,
	description string,
) *PermissionEntity {
	return &PermissionEntity{
		code:        code,
		name:        name,
		description: description,
	}
}
