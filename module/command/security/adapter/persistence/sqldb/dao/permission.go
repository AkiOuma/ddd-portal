package dao

import (
	"janus-portal/module/command/security/adapter/persistence/sqldb/model"
	"janus-portal/module/command/security/domain/entity"
	"janus-portal/module/command/security/domain/repo"

	"gorm.io/gorm"
)

// permission仓储实现
type PermissionDAO struct {
	db *gorm.DB
}

var _ repo.PermissionRepo = &PermissionDAO{}

func NewPermissionDAO(db *gorm.DB) *PermissionDAO {
	return &PermissionDAO{
		db: db,
	}
}

// 查询permission实体
func (dao PermissionDAO) FindPermissionCollection(
	userID,
	permissionCode,
	labelType string,
) (*entity.PermissionEntity, error) {
	filter := &model.PermissionModel{
		UserID:         userID,
		PermissionCode: permissionCode,
		LabelType:      labelType,
	}
	var result []model.PermissionModel
	if err := dao.db.Find(&result, filter).Error; err != nil {
		return nil, err
	}
	resources := make(map[string]struct{}, len(result))
	for _, v := range result {
		resources[v.LabelValue] = struct{}{}
	}
	return entity.NewPermissionEntity(
		userID,
		permissionCode,
		labelType,
		resources,
	), nil
}
