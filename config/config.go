package config

import (
	"log"

	aconf "github.com/AndreeJait/GO-ANDREE-UTILITIES/config"
)

type (
	EnvConfiguration struct {
		// - server
		ServerPort int `envconfig:"SERVER_PORT" default:"7010"`

		// - logger
		LoggerFileName   string `envconfig:"LOGGER_FILE_NAME" default:"./application.log" required:"true"`
		LoggerFormater   string `envconfig:"LOGGER_FORMATTER" default:"JSON" required:"true"`
		LoggerLevel      string `envconfig:"LOGGER_LEVEL" default:"INFO" required:"true"`
		LoggerMaxSize    int    `envconfig:"LOGGER_MAX_SIZE" default:"400" required:"true"`
		LoggerMaxBackups int    `envconfig:"LOGGER_MAX_BACKUPS" default:"0" required:"true"`
		LoggerMaxAge     int    `envconfig:"LOGGER_MAX_AGE" default:"7" required:"true"`
		LoggerCompress   bool   `envconfig:"LOGGER_COMPRESS" default:"true" required:"true"`

		// - mongo
		MongoDatabaseUri  string `envconfig:"MONGO_DATABASE_URI" default:"mongodb://0.0.0.0:27017" required:"true"`
		MongoDatabaseName string `envconfig:"MONGO_DATABASE_NAME" default:"delgong_story" required:"true"`

		// JWT
		ApplicationName string `envconfig:"APPLICATION_NAME" default:"DELGONG STORY BE" required:"true"`
		SecretKey       string `envconfig:"SECRET_KEY" default:"DELGONG HADEH" required:"true"`
	}
)

func NewEnvConfiguration() (*EnvConfiguration, error) {
	configuration := EnvConfiguration{}

	if err := aconf.NewFromEnv(&configuration); err != nil {
		log.Fatal(err)
	}

	return &configuration, nil
}
