package model

type RoleModel struct {
	ID         int `gorm:"column:id"`
	UserID     string
	RoleCode   string
	LabelType  string
	LabelValue string
}

func (RoleModel) TableName() string {
	return "role_user"
}
