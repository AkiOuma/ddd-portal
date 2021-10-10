package service

// 资源服务接口
type ResourceService interface {
	ValidateProjectCode(code string) bool
	ValidateSponsorCode(code string) bool
}
