package sqldb

import (
	"janus-portal/config/persistence"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewSqlDB(config *persistence.SqlDBConfig) *gorm.DB {
	dsn := buildDSN(config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: func() logger.Interface {
			return logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold:             time.Second,
					LogLevel:                  logger.Info,
					IgnoreRecordNotFoundError: true,
					Colorful:                  false,
				},
			)
		}(),
	})
	if err != nil {
		panic(err)
	}
	err = sqlDBMigrate(db)
	if err != nil {
		panic(err)
	}
	return db
}
