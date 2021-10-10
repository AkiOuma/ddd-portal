package application

import (
	"janus-portal/module/command/wxwork/domain/common"
	"janus-portal/module/command/wxwork/domain/repo"
	"janus-portal/module/command/wxwork/domain/valobj"
)

type WxworkApp struct {
	contactRepo repo.ContactRepo
}

func NewWxworkApp(
	contactRepo repo.ContactRepo,
) *WxworkApp {
	return &WxworkApp{
		contactRepo: contactRepo,
	}
}

// 获取通讯录中企业微信用户列表
func (app *WxworkApp) GetWxWorkUserList() ([]*valobj.User, error) {
	contact, err := app.contactRepo.NewContactEntity()
	if err != nil {
		return nil, common.ErrUnknownInternal
	}
	user, err := contact.GetWxWorkUserList()
	if err != nil {
		return nil, err
	}
	return user, nil
}
