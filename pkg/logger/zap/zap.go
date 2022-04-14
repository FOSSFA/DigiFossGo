package zap

import (
	_ "embed"
	"github.com/amirhossein-ka/DigiFossGo/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

type L struct {
	log *zap.SugaredLogger
}

func New(path string) logger.Logger {
	w := getLogWriter(path)
	e := getEncoder()

	core := zapcore.NewCore(e, w, zapcore.DebugLevel)
	newLog := zap.New(core, zap.AddCaller())

	return &L{log: newLog.Sugar()}
}

func getLogWriter(path string) zapcore.WriteSyncer {
	// go:embed
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	return zapcore.AddSync(file)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func (l *L) Info(field logger.LogField) {
	l.log.Infow(
		field.Message,
		"section", field.Section,
		"function", field.Function,
		"params", field.Params,
	)
}
func (l *L) Warn(field logger.LogField) {
	l.log.Warnw(
		field.Message,
		"section", field.Section,
		"function", field.Function,
		"params", field.Params,
	)
}

func (l *L) Error(field logger.LogField) {
	l.log.Errorw(
		field.Message,
		"section", field.Section,
		"function", field.Function,
		"params", field.Params,
	)
}
