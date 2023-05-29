package log

type (
	Config struct {
		// working mode dev/prod
		Mode string `yaml:"mode" validate:"required"`

		// format mode text/json
		LogFormat string `yaml:"log-format" validate:"required"`

		// log level debug/error
		LogLevel string `yaml:"log-level" validate:"required"`

		DateTimeFormat      string `yaml:"datetime-format"`
		UseTimestamp        bool   `yaml:"use-timestamp"`
		IncludeCallerMethod bool   `yaml:"include-caller-method"`

		// log to file
		OutputFilePath string `yaml:"output-file-path"`
	}
)
