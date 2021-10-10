package aggregate

import (
	"janus-portal/module/command/role_permission/domain/common"
	"janus-portal/module/command/role_permission/domain/entity"
	"time"
)

type SystemPermissionAgg struct {
	root       *entity.SystemEntity
	version    time.Time
	permission []*entity.PermissionEntity
}

func NewSystemPermissionAgg(
	root *entity.SystemEntity,
	permission []*entity.PermissionEntity,
) *SystemPermissionAgg {
	return &SystemPermissionAgg{
		root:       root,
		permission: permission,
	}
}

// 判断是否有新版本
func (agg *SystemPermissionAgg) isNewVersion(version time.Time) bool {
	return agg.version.Equal(version)
}

// 更新权限与新版本
func (agg *SystemPermissionAgg) UpdatePermission(
	version time.Time,
	permission []*entity.PermissionEntity,
) error {
	if !agg.isNewVersion(version) {
		return common.ErrVersionNotChange
	}
	agg.permission = permission
	return nil
}
