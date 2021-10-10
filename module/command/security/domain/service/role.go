package service

// 角色服务接口
type RoleService interface {
	ValidateRoleCode(code string) bool
}
