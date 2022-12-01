// 将zap日志写入到文件中

package main

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"net/http"
	"os"
	"time"
)

var (
	styleLogger  *zap.Logger
	sugarLogger2 *zap.SugaredLogger
)

func getLogWriter() zapcore.WriteSyncer {
	file, _ := rotatelogs.New(
		"./logs/system-%Y%m%d.log",
		rotatelogs.WithLinkName("./logs/system.log"),
		rotatelogs.WithMaxAge(7*time.Hour*24),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	ws := io.MultiWriter(file, os.Stdout)
	return zapcore.AddSync(ws)
}

func getJsonEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getConsoleEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func InitFileLogger() {
	writeSyncer := getLogWriter()
	encoder := getConsoleEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	styleLogger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	sugarLogger2 = styleLogger.Sugar()
}

func simpleHttpGet3(url string) {
	resp, err := http.Get(url)
	if err != nil {
		styleLogger.Error(
			"Error fetching url...",
			zap.String("url", url),
			zap.Error(err),
		)
	} else {
		sugarLogger2.Info(
			"Success...",
			zap.String("statusCode", resp.Status),
			zap.String("url", url),
		)
		resp.Body.Close()
	}
}

func main() {
	InitFileLogger()
	simpleHttpGet3("https://www.baidu.com")
}
