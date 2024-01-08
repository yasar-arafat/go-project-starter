package constants

const (
	ConfigFolderPath = "config"      // parent folder name of config file
	ConfigName       = "config"      // name of file without extension
	ConfigFileName   = "config.yaml" // name of config file with extension

	DBFolderPath         = "data"
	DBFileName           = "project-starter.db"
	DBMaxIdleConnections = 3

	LogFolderPath = "log"
	LogFileName   = "project-starter"

	ConfigType = "yaml"

	ServerPort = 8585

	LogLevel      = "INFO"
	LogMaxSize    = 5    // Maximum size in megabytes before log rotation
	LogMaxAge     = 3    // Maximum number of days to retain old log files
	LogMaxBackups = 28   // Maximum number of old log files to retain
	LogCompress   = true // Compress determines if the rotated log files should be compressed using gzip.
)
