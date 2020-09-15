package logger

import (
	"go.uber.org/zap"
)

var logger, _ = zap.NewDevelopment()
var Logger = logger.Sugar()
