package logger

import (
	kratoszap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"sync"
)

var logger *kratoszap.Logger
var once sync.Once

func Get() *kratoszap.Logger {
	once.Do(func() {
		config := zap.NewProductionConfig()
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		config.EncoderConfig.EncodeLevel = func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(strings.ToUpper(level.String()))
		}
		core := zapcore.NewCore(zapcore.NewConsoleEncoder(config.EncoderConfig), zapcore.Lock(os.Stderr), zap.DebugLevel)
		zlogger := zap.New(core).WithOptions()
		logger = kratoszap.NewLogger(zlogger)
	})
	return logger
}
