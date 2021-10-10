package sqldb

import (
	"janus-portal/config/persistence"

	"gorm.io/gorm"
)

func sqlDBMigrate(db *gorm.DB) error {
	return db.AutoMigrate()
}

func buildDSN(config *persistence.SqlDBConfig) string {
	return config.User() + ":" + config.Password() + "@tcp(" + config.Addr() + ")/" + config.Database() + "?parseTime=true"
}
