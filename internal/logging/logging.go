package logging

import (
	"log"
	"os"

	"github.com/christophersinclair/containment/internal/config"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

var (
	containmentLogger ContainmentLogger
	logLevel          int = INFO
	logLevelNames         = map[int]string{
		DEBUG: "DEBUG",
		INFO:  "INFO",
		WARN:  "WARN",
		ERROR: "ERROR",
		FATAL: "FATAL",
	}
)

type ContainmentLogger struct {
	logger *log.Logger
	level  int
}

func Setup(cfg *config.LoggingConfig) {
	var logger *log.Logger

	file, err := os.OpenFile(cfg.OutFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("WARN: Could not set up logging! %v", err)
		log.Printf("Logging will be to stdout instead of logfile...")
		logger = log.Default()
	} else {
		logger = log.New(file, "", log.LstdFlags)
	}

	containmentLogger.logger = logger

	setLogLevel(cfg.Level)
	containmentLogger.level = logLevel
}

func Get() *ContainmentLogger {
	return &containmentLogger
}

func (c *ContainmentLogger) LogMessage(level int, format string, v ...interface{}) {
	if level >= logLevel {
		c.logger.Printf("%s: ", logLevelNames[level])
		c.logger.Printf(format, v...)
	}
}

func setLogLevel(level string) {
	logLevel = 99
	if level == "" {
		logLevel = WARN
	}

	if level == "DEBUG" {
		logLevel = DEBUG
	}

	if level == "INFO" {
		logLevel = INFO
	}

	if level == "WARN" {
		logLevel = WARN
	}

	if level == "ERROR" {
		logLevel = ERROR
	}

	if level == "FATAL" {
		logLevel = FATAL
	}

	if logLevel == 99 { // nothing was set
		logLevel = WARN
	}
}
