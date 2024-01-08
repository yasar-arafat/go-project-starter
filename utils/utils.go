package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/labstack/gommon/log"
)

func CreateFolder(folderPath string) error {
	// Check if the folder already exists
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// Create the folder if it doesn't exist
		err := os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			return err
		}
		fmt.Printf("Folder created: %s\n", folderPath)
	} else if err != nil {
		// Handle other errors that may have occurred during the Stat operation
		return err
	} else {
		fmt.Printf("Folder already exists: %s\n", folderPath)
	}
	return nil
}

// ParseLogLevel parses a string into a Log.Lvl
func ParseLogLevel(logLvlStr string) (log.Lvl, error) {
	var err error
	var logLvl log.Lvl
	switch strings.ToUpper(logLvlStr) {
	case "DEBUG":
		logLvl = log.DEBUG
	case "INFO":
		logLvl = log.INFO
	case "WARN":
		logLvl = log.WARN
	case "ERROR":
		logLvl = log.ERROR
	case "OFF":
		logLvl = log.OFF
	default:
		logLvl = log.INFO
		err = errors.New("invalid log level")
	}
	return logLvl, err
}
