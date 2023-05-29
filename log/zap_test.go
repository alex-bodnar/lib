package log

import (
	"testing"
)

func Test_zapLog_Errorw(t *testing.T) {
	cfg := Config{
		Mode:                "dev",
		LogFormat:           "json",
		LogLevel:            "debug",
		IncludeCallerMethod: true,
	}

	logger := InitLogger(cfg, map[string]string{
		"field": "value",
	})

	logger.Debug("Blas", "test", "t")
	logger.Info("Blas", "test", "t")
	logger.Warning("Blas", "test", "t")
}

func Test_zapLog_NOT_IncludeCallerMethod(t *testing.T) {
	cfg := Config{
		Mode:                "dev",
		LogFormat:           "text",
		LogLevel:            "debug",
		IncludeCallerMethod: false,
	}

	logger := InitLogger(cfg, map[string]string{
		"field": "value",
	})

	logger.Debug("Blas", "test", "t")
	logger.Info("Blas", "test", "t")
	logger.Warning("Blas", "test", "t")
}
