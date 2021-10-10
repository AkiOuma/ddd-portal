package valobj

import (
	"janus-portal/module/command/security/domain/common"
	"janus-portal/module/command/security/domain/service"
)

type ResourceLabel struct {
	labelType string
	value     string
}

func NewResourceObject(
	labelType string,
	value string,
	resourceService service.ResourceService,
) (*ResourceLabel, error) {
	vo := &ResourceLabel{}
	if err := vo.setLabelType(labelType); err != nil {
		return nil, err
	}
	if err := vo.setValue(value, resourceService); err != nil {
		return nil, err
	}
	return vo, nil
}

// 校验标签类型
func (vo *ResourceLabel) validateLabelType(labelType string) error {
	for _, v := range [3]string{"sponsor", "project", "*"} {
		if labelType == v {
			return nil
		}
	}
	return common.ErrInvalidResourceType
}

// 设置标签类型
func (vo *ResourceLabel) setLabelType(labelType string) error {
	if err := vo.validateLabelType(labelType); err != nil {
		return err
	}
	return nil
}

// 获取标签类型
func (vo *ResourceLabel) LabelType() string {
	return vo.labelType
}

// 校验标签值
func (vo *ResourceLabel) validateValue(
	value string,
	resourceService service.ResourceService,
) error {
	switch {
	case value == "*" && vo.labelType == "*":
		return nil
	case vo.labelType == "sponsor":
		if !resourceService.ValidateSponsorCode(value) {
			return common.ErrInvalidResourceValue
		}
	case vo.labelType == "project":
		if !resourceService.ValidateProjectCode(value) {
			return common.ErrInvalidResourceValue
		}
	}
	return nil
}

// 设置标签值
func (vo *ResourceLabel) setValue(
	value string,
	resourceService service.ResourceService,
) error {
	if err := vo.validateValue(value, resourceService); err != nil {
		return err
	}
	vo.value = value
	return nil
}

// 获取标签值
func (vo *ResourceLabel) Value() string {
	return vo.value
}
