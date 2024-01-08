package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/yasar-arafat/go-project-starter/constants"
	"github.com/yasar-arafat/go-project-starter/utils"
)

// Config represents the configuration structure
type Config struct {
	// Server configuration
	Server struct {
		Port string // Port on which the server will listen
	}

	// Database configuration
	Database struct {
		FileName           string // Filename for the database
		MaxIdleConnections int    // MaxIdleConnections is the maximum number of idle connections to the database.
	}

	// Log configuration
	Log struct {
		Level      string // Log level (e.g., DEBUG, INFO, WARN, ERROR)
		FileName   string `json:"filename" yaml:"filename"`     // Filename for log output
		MaxSize    int    `json:"maxsize" yaml:"maxsize"`       // Maximum size in megabytes before log rotation
		MaxAge     int    `json:"maxage" yaml:"maxage"`         // Maximum number of days to retain old log files
		MaxBackups int    `json:"maxbackups" yaml:"maxbackups"` // Maximum number of old log files to retain
		Compress   bool   `json:"compress" yaml:"compress"`     // Compress determines if the rotated log files should be compressed using gzip.
	}

	// Add other configuration fields as needed
}

// Read reads the configuration from the file, creates a new configuration file with default values if it doesn't exist,
// and returns a Config struct.
func Read() (*Config, error) {
	// Set the file path for the configuration file
	configFilePath := filepath.Join(constants.ConfigFolderPath, constants.ConfigFileName)

	// Set default values if the configuration file doesn't exist
	SetDefaultValues()

	// Check if the configuration file exists
	_, err := os.Stat(configFilePath)
	if os.IsNotExist(err) {
		// Create the configuration directory if it doesn't exist
		if err := utils.CreateFolder(constants.ConfigFolderPath); err != nil {
			return nil, fmt.Errorf("failed to create config directory")
		}

		// Write the default configuration to a new file
		err = viper.WriteConfigAs(configFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create config file: %v", err)
		}
		fmt.Println("Config file created with default values.")
	} else if err != nil {
		return nil, err
	} else {
		// Read the config file
		viper.AddConfigPath(constants.ConfigFolderPath)
		viper.SetConfigName(constants.ConfigName)
		viper.SetConfigType(constants.ConfigType)
		err = viper.ReadInConfig()
		if err != nil {
			return nil, fmt.Errorf("failed to read config file: %v", err)
		}
	}

	// Unmarshal the config values into a struct
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return &config, nil
}

// SetDefaultValues sets the default values for the configuration using Viper.
func SetDefaultValues() {
	viper.SetDefault("Server.Port", constants.ServerPort)
	viper.SetDefault("Database.MaxIdleConnections", constants.DBMaxIdleConnections)
	viper.SetDefault("Log.Compress", constants.LogCompress)
	viper.SetDefault("Log.FileName", constants.LogFileName)
	viper.SetDefault("Log.Level", constants.LogLevel)
	viper.SetDefault("Log.MaxSize", constants.LogMaxSize)
	viper.SetDefault("Log.MaxAge", constants.LogMaxAge)
	viper.SetDefault("Log.MaxBackups", constants.LogMaxBackups)
	viper.SetDefault("Log.Compress", constants.LogCompress)

}
