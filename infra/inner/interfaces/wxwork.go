package interfaces

import "janus-portal/types/inner"

type Wxwork interface {
	GetWxWorkUserList() ([]*inner.UserDetail, error)
}
