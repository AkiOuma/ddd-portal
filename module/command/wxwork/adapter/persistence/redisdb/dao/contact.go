package dao

import (
	"janus-portal/module/command/wxwork/domain/entity"
	"janus-portal/module/command/wxwork/domain/repo"

	"github.com/go-redis/redis/v8"
)

type ContactDAO struct {
	rdb *redis.Client
}

func NewContactDAO(rdb *redis.Client) *ContactDAO {
	return &ContactDAO{
		rdb: rdb,
	}
}

var _ repo.ContactRepo = &ContactDAO{}

func (dao *ContactDAO) NewContactEntity() (*entity.ContactEntity, error) {
	return entity.NewContactEntity(""), nil
}
