package repo

import "janus-portal/module/command/security/domain/entity"

type RoleRepo interface {
	SaveRoleCollection(role *entity.RoleEntity) error
}
