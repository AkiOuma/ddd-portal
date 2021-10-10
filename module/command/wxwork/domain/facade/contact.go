package facade

import (
	"encoding/json"
	"janus-portal/module/command/wxwork/domain/common"
	"janus-portal/module/command/wxwork/domain/valobj"
	"net/http"
)

func GetWxWorkUserList(url string) ([]*valobj.User, error) {
	// get http result
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// decode
	result := common.SyncUserResponse{}
	result.UserList = make([]common.WxworkUserInfo, 0)
	json.NewDecoder(resp.Body).Decode(&result)

	// transfer into value object
	user := make([]*valobj.User, len(result.UserList))
	for k, v := range result.UserList {
		user[k] = valobj.NewUser(
			v.UserID,
			v.Username,
			v.Avatar,
			v.Email,
			v.Mobile,
		)
	}
	return user, nil
}
