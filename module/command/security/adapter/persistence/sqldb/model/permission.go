package model

type PermissionModel struct {
	UserID         string `gorm:"column:userid"`
	PermissionCode string `gorm:"column:permission_code"`
	LabelType      string `gorm:"column:label_type"`
	LabelValue     string `gorm:"column:label_value"`
}

func (PermissionModel) TableName() string {
	return "permission_collection"
}
