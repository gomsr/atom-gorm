package gormc

import (
	"fmt"
	"gorm.io/gorm/logger"
	"log"
)

const (
	Console = "console"
	Zap     = "zap"
	GoZero  = "go-zero"
)

type writer struct {
	logger.Writer
	DbType  string
	LogType string
}

// NewWriter writer 构造函数
func NewWriter(logType, dbType string, w logger.Writer) *writer {
	return &writer{Writer: w, DbType: dbType, LogType: logType}
}

// Printf 格式化打印日志
func (w *writer) Printf(message string, data ...interface{}) {
	log.Printf(message+"\n", data...)
	fmt.Printf(message+"\n", data...)
	w.Writer.Printf(message, data...)
}
