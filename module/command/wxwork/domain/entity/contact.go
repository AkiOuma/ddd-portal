package entity

import (
	"janus-portal/module/command/wxwork/domain/common"
	"janus-portal/module/command/wxwork/domain/facade"
	"janus-portal/module/command/wxwork/domain/valobj"
)

type ContactEntity struct {
	accessToken string
	url         string
}

func NewContactEntity(token string) *ContactEntity {
	return &ContactEntity{
		accessToken: token,
	}
}

// 获取通讯录中企业微信用户列表（参考内容：https://work.weixin.qq.com/api/doc/90000/90135/90201）
//
// @return: 企业微信用户列表
func (entity *ContactEntity) GetWxWorkUserList() ([]*valobj.User, error) {
	user, err := facade.GetWxWorkUserList(entity.buildContactUrl())
	if err != nil {
		return nil, common.ErrSyncUser
	}
	return user, nil
}

// 构建通讯录接口url
func (entity *ContactEntity) buildContactUrl() string {
	return entity.url + "?" + "access_token=" + entity.accessToken + "&department_id=1&fetch_child=1"
}
