package dao

import (
	"janus-portal/module/command/security/adapter/persistence/sqldb/model"
	"janus-portal/module/command/security/domain/entity"
	"janus-portal/module/command/security/domain/repo"

	"gorm.io/gorm"
)

type RoleDAO struct {
	db *gorm.DB
}

var _ repo.RoleRepo = &RoleDAO{}

func NewRoleDAO(db *gorm.DB) *RoleDAO {
	return &RoleDAO{
		db: db,
	}
}

// 持久化一个角色实体
func (dao *RoleDAO) SaveRoleCollection(role *entity.RoleEntity) error {
	row := &model.RoleModel{
		UserID:     role.UserID(),
		RoleCode:   role.RoleCode(),
		LabelType:  role.LabelType(),
		LabelValue: role.LabelValue(),
	}
	return dao.db.Create(row).Error
}
