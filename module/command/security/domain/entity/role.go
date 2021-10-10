package entity

import "janus-portal/module/command/security/domain/valobj"

type RoleEntity struct {
	userID        string
	roleCode      *valobj.RoleCode
	resourceLabel *valobj.ResourceLabel
}

func NewRoleEntity(
	userID string,
	roleCode *valobj.RoleCode,
	resourceLabel *valobj.ResourceLabel,
) *RoleEntity {
	return &RoleEntity{
		userID:        userID,
		roleCode:      roleCode,
		resourceLabel: resourceLabel,
	}
}

func (entity *RoleEntity) UserID() string {
	return entity.userID
}

func (entity *RoleEntity) RoleCode() string {
	return entity.roleCode.Code()
}

func (entity *RoleEntity) LabelType() string {
	return entity.resourceLabel.LabelType()
}
func (entity *RoleEntity) LabelValue() string {
	return entity.resourceLabel.Value()
}
