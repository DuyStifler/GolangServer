package utils

import (
	"DuyStifler/GolangServer/models"
)

type Logger struct {
	errorLogPath string
	infoLogPath  string
}

func NewLogger(errorPath, infoPath string) (*Logger, error) {
	 serverLogger := &Logger{errorLogPath: errorPath, infoLogPath: infoPath}
	 return serverLogger, nil
}

func (l *Logger) LogError(logError models.Error) {

}

func (l *Logger) LogInfo(logError models.Error) {
	
}