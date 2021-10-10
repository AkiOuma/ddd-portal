package valobj

import (
	"janus-portal/module/command/security/domain/common"
	"janus-portal/module/command/security/domain/service"
)

type RoleCode struct {
	code string
}

func NewRoleCode(
	code string,
	roleService service.RoleService,
) (*RoleCode, error) {
	vo := &RoleCode{}
	if err := vo.setCode(code, roleService); err != nil {
		return nil, err
	}
	return vo, nil
}

// 校验角色代码
func (vo *RoleCode) validateCode(
	code string,
	roleService service.RoleService,
) error {
	if !roleService.ValidateRoleCode(code) {
		return common.ErrInvalidRoleCode
	}
	return nil
}

// 设置角色代码
func (vo *RoleCode) setCode(
	code string,
	roleService service.RoleService,
) error {
	if err := vo.validateCode(code, roleService); err != nil {
		return err
	}
	vo.code = code
	return nil
}

// 获取角色代码
func (vo *RoleCode) Code() string {
	return vo.code
}
