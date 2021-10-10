package repo

import "janus-portal/module/command/wxwork/domain/entity"

type ContactRepo interface {
	NewContactEntity() (*entity.ContactEntity, error)
}
