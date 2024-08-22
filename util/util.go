package util

import (
	"flag"
	"net/url"
	"os"

	"github.com/ankitg1706/zoopone/model"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var Logger logrus.Logger

func init() {

	Logger = *logrus.New()

	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Logger.SetOutput(os.Stdout)
}

func SetLogger() {

	logLevel := flag.String(model.LogLevel, model.LogLevelInfo, "log-level (debug , error , info, warning )")
	flag.Parse()
	switch *logLevel {
	case model.LegLevelDebug:
		Logger.SetLevel(logrus.DebugLevel)
	case model.LogLevelError:
		Logger.SetLevel(logrus.ErrorLevel)
	case model.LogLevelWarning:
		Logger.SetLevel(logrus.WarnLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}

}

func Log(logLevel, packageLevel, functionName string, message, paramerer interface{}) {
	switch logLevel {
	case model.LegLevelDebug:
		if paramerer != nil {
			Logger.Debugf("packageLevel: %s, functionName: %s, message: %v, paramerer: %v\n", packageLevel, functionName, message, paramerer)
		} else {
			Logger.Debugf("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)
		}
	case model.LogLevelError:
		if paramerer != nil {
			Logger.Errorf("packageLevel: %s, functionName: %s, message: %v, paramerer: %v\n", packageLevel, functionName, message, paramerer)
		} else {
			Logger.Errorf("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)
		}
	case model.LogLevelWarning:
		if paramerer != nil {
			Logger.Warnf("packageLevel: %s, functionName: %s, message: %v, paramerer: %v\n", packageLevel, functionName, message, paramerer)
		} else {
			Logger.Warnf("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)
		}
	default:
		if paramerer != nil {
			Logger.Infof("packageLevel: %s, functionName: %s, message: %v, paramerer: %v\n", packageLevel, functionName, message, paramerer)
		} else {
			Logger.Infof("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)
		}
	}
}

// ConvertQueryParams converts url.Values to map[string]interface{}
func ConvertQueryParams(queryParams url.Values) map[string]interface{} {
	result := make(map[string]interface{})

	for key, values := range queryParams {
		if key == "id" {
			uuid, _ := uuid.Parse(values[0])
			result[key] = uuid
			continue
		}
		if len(values) == 1 {
			result[key] = values[0] // single value, add as string
		} else {
			result[key] = values // multiple values, add as []string
		}
	}

	return result
}
