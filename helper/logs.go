package helper

import (
	"GolangBackend/constants"
	"GolangBackend/internal/global"
	"os"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)


func InitLogger(isWriteLogsToFile bool) {
	global.Logger = log.New()

	global.Logger.SetFormatter(
		&log.TextFormatter{
			TimestampFormat: time.RFC3339,
			FullTimestamp:   true,
			ForceColors:     true,
		},
	)

	global.Logger.SetLevel(log.DebugLevel)

	if isWriteLogsToFile {
		file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, constants.OS_READ_WRITE_FILE_PERMISSION)
		if err != nil {
			global.Logger.Error("Failed to write logs to file")
		} else {
			global.Logger.SetOutput(file)
		}
	}

	global.Logger.Infoln("Init Logger successfully")
}

func LogError(err error, msg string) {
	global.Logger.Errorf("%s, %v", msg, errors.WithStack(err))
}

func LogWarn(msg string, args ...any) {
	global.Logger.Warnf(msg, args...)
}

func LogInfo(msg string, args ...any) {
	global.Logger.Infof(msg, args...)
}
