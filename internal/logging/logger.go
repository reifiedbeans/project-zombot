package logging

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func NewLogger() *zap.SugaredLogger {
	config := zap.NewProductionConfig()
	config.Encoding = "console"
	config.EncoderConfig = zap.NewDevelopmentEncoderConfig()
	logger, err := config.Build()
	if err != nil {
		err = errors.Wrap(err, "Could not create logger")
		panic(err)
	}
	return logger.Sugar()
}
