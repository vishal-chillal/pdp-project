package logger_test

import (
	"testing"
	"logger"
	)
func Test_Init(t *testing.T) {
	//t.Fail()
	
	logConfig := &logger.LoggerConfig{}
	if !logger.Init(logConfig){
		t.Fail()
	}
	
}