package repo

import (
	"janus-portal/module/command/role_permission/domain/aggregate"
	"janus-portal/module/command/role_permission/domain/entity"
)

type SystemRepo interface {
	SaveSystemInfo(*entity.SystemEntity) error
	SaveSystemPermission(*aggregate.SystemPermissionAgg) error
}
