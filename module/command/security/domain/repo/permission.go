package repo

import "janus-portal/module/command/security/domain/entity"

type PermissionRepo interface {
	FindPermissionCollection(
		userID,
		permissionCode,
		labelType string,
	) (*entity.PermissionEntity, error)
}
