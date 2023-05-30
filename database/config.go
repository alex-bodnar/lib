package database

import (
	"time"
)

type (
	// Config - configuration for postgreSQL database
	Config struct {
		ConnectionString   string        `yaml:"connection-string" validate:"required"`
		ConnMaxIdleNum     int           `yaml:"conn-max-idle-num"`
		ConnMaxOpenNum     int           `yaml:"conn-max-open-num"`
		Dialect            string        `yaml:"dialect" validate:"required"`
		Driver             string        `yaml:"driver" validate:"required"`
		MaxRetries         int           `yaml:"max-retries" validate:"required"`
		ConnMaxLifetime    time.Duration `yaml:"conn-max-lifetime"`
		RetryDelay         time.Duration `yaml:"retry-delay" validate:"required"`
		QueryTimeout       time.Duration `yaml:"query-timeout"`
		AutoMigrate        bool          `yaml:"auto-migrate"`
		MigrationDirectory string        `yaml:"migration-directory" validate:"required"`
	}
)
