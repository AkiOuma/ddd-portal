package wxwork

import "os"

type ContactConfig struct {
	corpsecret string
	url        string
}

func NewContactConfig() *ContactConfig {
	return &ContactConfig{
		corpsecret: os.Getenv("wx_corpsecret_contact"),
		url:        os.Getenv("wx_contact_url"),
	}
}

func (config *ContactConfig) CorpSecret() string {
	return config.corpsecret
}

func (config *ContactConfig) Url() string {
	return config.url
}
