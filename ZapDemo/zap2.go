package main

import (
	"go.uber.org/zap"
	"net/http"
)

var (
	sugarLogger *zap.SugaredLogger
)

func InitSugarLogger() {
	logger, _ := zap.NewProduction()
	sugarLogger = logger.Sugar()
}

func simpleHttpGet2(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		resp.Body.Close()
	}
}

func main() {
	InitSugarLogger()
	simpleHttpGet2("baidu.com")
}
