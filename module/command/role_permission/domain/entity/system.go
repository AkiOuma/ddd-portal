package entity

import (
	"janus-portal/module/command/role_permission/domain/valobj"
)

type SystemEntity struct {
	code        *valobj.Code
	name        string
	description string
}

func NewSystemEntity(
	code *valobj.Code,
	name string,
	description string,
) *SystemEntity {
	return &SystemEntity{
		code:        code,
		name:        name,
		description: description,
	}
}

// 更新系统名称
func (entity *SystemEntity) UpdateName(name string) {
	entity.name = name
}

// 更新系统描述
func (entity *SystemEntity) UpdateDescription(content string) {
	entity.description = content
}
