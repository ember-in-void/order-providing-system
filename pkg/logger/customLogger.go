package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

// CustomLogger для логирования с разделением по уровням
type CustomLogger struct {
	infoLogger *log.Logger
	warnLogger *log.Logger
	errLogger  *log.Logger
	mu         sync.RWMutex
}

func NewCustomLogger() (*CustomLogger, error) {
	flags := log.Ldate | log.Ltime | log.Lshortfile

	// Убедимся, что директория существует
	err := os.MkdirAll("logs", 0o755)
	if err != nil {
		return nil, fmt.Errorf("failed to create logs directory: %w", err)
	}

	// Открываем файлы для логирования
	fileInfo, err := os.OpenFile("logs/info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		return nil, fmt.Errorf("failed to open info log file: %w", err)
	}
	fileWarn, err := os.OpenFile("logs/warning.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		return nil, fmt.Errorf("failed to open warning log file: %w", err)
	}
	fileErr, err := os.OpenFile("logs/error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		return nil, fmt.Errorf("failed to open error log file: %w", err)
	}

	Logger := &CustomLogger{
		infoLogger: log.New(fileInfo, "INFO: ", flags),
		warnLogger: log.New(fileWarn, "WARN: ", flags),
		errLogger:  log.New(fileErr, "ERROR: ", flags),
	}

	return Logger, nil
}

func (l *CustomLogger) Info(msg ...interface{}) {
	if l == nil || l.infoLogger == nil {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	l.infoLogger.Println(msg...)
}

func (l *CustomLogger) Warn(msg ...interface{}) {
	if l == nil || l.warnLogger == nil {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	l.warnLogger.Println(msg...)
}

func (l *CustomLogger) Error(msg ...interface{}) {
	if l == nil || l.errLogger == nil {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	l.errLogger.Println(msg...)
}
