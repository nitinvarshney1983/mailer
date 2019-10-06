package logging

import (
	"io"
	"os"

	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

// LogFileConfigs to set up logs file location
type LogFileConfigs struct {
	logPath  string
	fileName string
}

// LoggerConfigs to set log level and to point log file configurations
type LoggerConfigs struct {
	level          string
	logFileConfigs LogFileConfigs
}

//Logger struct, it would be used in application
type Logger struct {
	logger *log.Logger
}

//Log interface to define the method used by Logger
type Log interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Fatel(args ...interface{})
	Panic(args ...interface{})
	DebugWithFields(fields map[string]interface{}, args ...interface{})
	InfoWithFields(fields map[string]interface{}, args ...interface{})
	WarnWithFields(fields map[string]interface{}, args ...interface{})
	ErrorWithFields(fields map[string]interface{}, args ...interface{})
	FatalWithFields(fields map[string]interface{}, args ...interface{})
	PanicWithFields(fields map[string]interface{}, args ...interface{})
}

var instance *Logger

func init() {
	instance = &Logger{
		logger: log.New(),
	}
	instance.logger.SetLevel(log.InfoLevel)
}

// Setup to set up logging level and logging file location
func Setup(configurations interface{}) {
	cfg, ok := configurations.(map[string]interface{})
	if !ok {
		log.Error("not able to read configurations properly")
		return
	}
	if envLevel, found := os.LookupEnv("LOG_LEVEL"); found {
		cfg["level"] = envLevel
	}
	if cfg["level"] != nil {
		instance.logger.SetLevel(internalLogLevel((cfg["level"].(string))))
	}

	if cfg["logfileconfigs"] != nil {
		logfileConfigs, _ := cfg["logfileconfigs"].(map[string]interface{})
		logFileName := logfileConfigs["filename"]
		logFilePath := logfileConfigs["logpath"]
		lumberjackLogrotate := &lumberjack.Logger{
			Filename:   logFilePath.(string) + "/" + logFileName.(string),
			MaxSize:    1,  // Max megabytes before log is rotated
			MaxBackups: 90, // Max number of old log files to keep
			MaxAge:     60, // Max number of days to retain log files
			Compress:   true,
		}
		logMultiWriter := io.MultiWriter(os.Stdout, lumberjackLogrotate)
		instance.logger.SetOutput(logMultiWriter)
		instance.logger.SetFormatter(&log.JSONFormatter{})
	}

}

func internalLogLevel(level string) (internalLevel log.Level) {
	switch level {
	case "PANIC":
		internalLevel = log.PanicLevel
	case "FATAL":
		internalLevel = log.FatalLevel
	case "ERROR":
		internalLevel = log.ErrorLevel
	case "INFO":
		internalLevel = log.InfoLevel
	case "WARN":
		internalLevel = log.WarnLevel
	case "DEBUG":
		internalLevel = log.DebugLevel
	default:
		internalLevel = log.InfoLevel
	}
	return

}

//Panic logger
func Panic(args ...interface{}) {
	instance.logger.Panic(args...)
}

//Fatal logger
func Fatal(args ...interface{}) {
	instance.logger.Fatal(args...)
}

//Warn logger
func Warn(args ...interface{}) {
	instance.logger.Warn(args...)
}

//Info logger
func Info(args ...interface{}) {
	instance.logger.Info(args...)
}

//Debug logger
func Debug(args ...interface{}) {
	instance.logger.Debug(args...)
}

//DebugWithFields logger
func DebugWithFields(fields map[string]interface{}, args ...interface{}) {
	instance.logger.Debug(fields, args)
}

//InfoWithFields logger
func InfoWithFields(fields map[string]interface{}, args ...interface{}) {
	instance.logger.Info(fields, args)
}

//WarnWithFields Logger
func WarnWithFields(fields map[string]interface{}, args ...interface{}) {
	instance.logger.Warn(fields, args)
}

//ErrorWithFields logger
func ErrorWithFields(fields map[string]interface{}, args ...interface{}) {
	instance.logger.Error(fields, args)
}

//FatalWithFields logger
func FatalWithFields(fields map[string]interface{}, args ...interface{}) {
	instance.logger.Fatal(fields, args)
}

//PanicWithFields logger
func PanicWithFields(fields map[string]interface{}, args ...interface{}) {
	instance.logger.Panic(fields, args)
}
