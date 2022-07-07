package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

var sugar *zap.SugaredLogger

func Init() {

	// zap logger
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(timeFormat)

	file, _ := os.Create("logs/gloves.log")
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(file), zapcore.AddSync(os.Stdout)),
		zapcore.DebugLevel,
	)

	zl := zap.New(core)
	defer zl.Sync()
	zap.ReplaceGlobals(zl)

	sugar = zl.Sugar()
}

func Debug(v ...interface{}) {
	sugar.Debug(v...)
}

func Debugf(format string, args ...interface{}) {
	sugar.Debugf(format, args...)
}

func Info(v ...interface{}) {
	sugar.Info(v...)
}

func Infof(format string, args ...interface{}) {
	sugar.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	sugar.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	sugar.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	sugar.Fatalf(format, args...)
}
