package application

import (
	"janus-portal/module/command/security/domain/entity"
	"janus-portal/module/command/security/domain/repo"
	"janus-portal/module/command/security/domain/service"
	"janus-portal/module/command/security/domain/valobj"
)

type SecurityApp struct {
	permissionRepo  repo.PermissionRepo
	roleRepo        repo.RoleRepo
	resourceService service.ResourceService
	roleService     service.RoleService
}

func NewSecurityApp(
	permissionRepo repo.PermissionRepo,
	roleRepo repo.RoleRepo,
	resourceService service.ResourceService,
	roleService service.RoleService,
) *SecurityApp {
	return &SecurityApp{
		permissionRepo:  permissionRepo,
		roleRepo:        roleRepo,
		resourceService: resourceService,
		roleService:     roleService,
	}
}

// 校验权限组合
func (app *SecurityApp) ValidatePermission(
	userID,
	permissionCode,
	labelType,
	labelValue string,
) bool {
	collection, err := app.permissionRepo.FindPermissionCollection(
		userID,
		permissionCode,
		labelType,
	)
	if err != nil {
		return false
	}
	return !collection.HasResource(labelValue)
}

// 查询持有资源
func (app *SecurityApp) GetResourceList(
	userID,
	permissionCode,
	labelType string,
) ([]string, error) {
	collection, err := app.permissionRepo.FindPermissionCollection(
		userID,
		permissionCode,
		labelType,
	)
	if err != nil {
		return nil, err
	}
	return collection.Resources(), nil
}

// 用户绑定角色
func (app *SecurityApp) BindRole(
	userID,
	code,
	labelType,
	labelValue string,
) error {
	roleCode, err := valobj.NewRoleCode(code, app.roleService)
	if err != nil {
		return err
	}
	resource, err := valobj.NewResourceObject(
		labelType,
		labelValue,
		app.resourceService,
	)
	if err != nil {
		return err
	}
	role := entity.NewRoleEntity(
		userID,
		roleCode,
		resource,
	)
	return app.roleRepo.SaveRoleCollection(role)
}
