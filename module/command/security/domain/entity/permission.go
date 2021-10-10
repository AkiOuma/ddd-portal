package entity

type PermissionEntity struct {
	userID         string
	permissionCode string
	resourceType   string
	resources      map[string]struct{}
}

func NewPermissionEntity(
	userID string,
	permissionCode string,
	resourceType string,
	resources map[string]struct{},
) *PermissionEntity {
	return &PermissionEntity{
		userID:         userID,
		permissionCode: permissionCode,
		resourceType:   resourceType,
		resources:      resources,
	}
}

// 是否持有特定资源
func (entity *PermissionEntity) HasResource(resourceCode string) bool {
	_, ok := entity.resources[resourceCode]
	return ok
}

// 获取资源列表
func (entity *PermissionEntity) Resources() []string {
	resources := make([]string, len(entity.resources))
	index := 0
	for k := range entity.resources {
		resources[index] = k
		index++
	}
	return resources
}
