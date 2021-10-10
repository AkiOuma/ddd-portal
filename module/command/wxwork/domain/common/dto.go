package common

type WxworkUserInfo struct {
	UserID   string `json:"userid"`
	Username string `json:"name"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
}

type SyncUserResponse struct {
	ErrCode    int              `json:"errcode"`
	ErrMessage string           `json:"errmsg"`
	UserList   []WxworkUserInfo `json:"userlist"`
}
