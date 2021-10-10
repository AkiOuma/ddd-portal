package inner

import (
	"janus-portal/module/command/wxwork/application"
	"janus-portal/types/inner"
)

type WxworkInnerService struct {
	app *application.WxworkApp
}

func NewWxworkInnerService(app *application.WxworkApp) *WxworkInnerService {
	return &WxworkInnerService{
		app: app,
	}
}

// 获取通讯录中企业微信用户列表
func (service *WxworkInnerService) GetWxWorkUserList() ([]*inner.UserDetail, error) {
	users, err := service.app.GetWxWorkUserList()
	if err != nil {
		return nil, err
	}
	userList := make([]*inner.UserDetail, len(users))
	for k, v := range users {
		userList[k] = &inner.UserDetail{
			UserID:   v.UserID(),
			Username: v.Username(),
			Avatar:   v.Avatar(),
			Email:    v.Email(),
			Mobile:   v.Mobile(),
		}
	}
	return userList, nil
}
