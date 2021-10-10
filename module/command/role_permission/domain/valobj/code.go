package valobj

import (
	"janus-portal/module/command/role_permission/domain/common"
	"regexp"
)

type Code struct {
	value string
}

func NewCode(value string) (*Code, error) {
	code := &Code{}
	if err := code.setCode(value); err != nil {
		return nil, err
	}
	return code, nil
}

// 校验代码有效性
func (Code) validateCode(value string) bool {
	pattern := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return pattern.MatchString(value)
}

// 设置代码
func (vo *Code) setCode(value string) error {
	if !vo.validateCode(value) {
		return common.ErrInvalidCode
	}
	vo.value = value
	return nil
}

// 获取代码
func (vo *Code) Value() string {
	return vo.value
}
