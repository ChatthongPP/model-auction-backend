package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host     string `env:"THIRDPARTY_POS_HOST"`
	Username string `env:"THIRDPARTY_POS_USERNAME"`
	Password string `env:"THIRDPARTY_POS_PASSWORD"`
	Database string `env:"THIRDPARTY_POS_DATABASE"`
	Port     int    `env:"THIRDPARTY_POS_PORT"`
}

// BuildDSN func
func buildDSN(config *Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.Database,
	)
}

func ConnectDB(dbConf *Config) *gorm.DB {
	dsn := buildDSN(dbConf)
	log.Printf("dsn: %s\n", dsn)
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// DisableForeignKeyConstraintWhenMigrating: true,
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		panic(`fatal error: cannot connect to database`)
	}

	return dbConn
}
