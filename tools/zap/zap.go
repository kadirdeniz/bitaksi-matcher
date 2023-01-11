package zap

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"runtime"
)

var Logger *zap.Logger

func InitilizeLogger() {
	zapConfig := zap.NewProductionConfig()

	fileEncoder := zapcore.NewJSONEncoder(zapConfig.EncoderConfig)

	_, b, _, _ := runtime.Caller(0)
	docsBasePath := filepath.Join(filepath.Dir(b), "./../../docs")

	logFile, err := os.Create(filepath.Join(docsBasePath, "log.json"))
	if err != nil {
		fmt.Println(err)
	}
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(fileEncoder, zapcore.Lock(os.Stdout), defaultLogLevel),
	)

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}
