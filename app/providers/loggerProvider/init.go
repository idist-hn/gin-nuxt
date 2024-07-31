package loggerProvider

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"idist-core/app/providers/configProvider"
	"os"
	"path/filepath"
)

var Logger *zap.SugaredLogger
var LoggerRequest *zap.SugaredLogger

func GetLogger() *zap.SugaredLogger {
	if Logger == nil {
		Init()
	}
	return Logger
}
func Init() {
	fmt.Println("------------------------------------------------------------")
	//var err error
	cf := configProvider.GetConfig()

	if dir, err := os.Getwd(); err != nil {
		panic(err)
	} else if fi, err := os.Stat(filepath.Join(dir, "logs")); os.IsNotExist(err) || !fi.IsDir() {
		_ = os.MkdirAll(filepath.Join(dir, "logs"), 0755)
	}
	cfg := zap.NewDevelopmentConfig()
	switch cf.GetString("log.level") {
	case "debug":
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		break
	case "info":
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		break
	case "warning":
		cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
		break
	case "error":
		cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
		break
	case "dpanic":
		cfg.Level = zap.NewAtomicLevelAt(zap.DPanicLevel)
		break
	case "panic":
		cfg.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
		break
	}

	cfg.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "data",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	LoggerDefaultInit(cfg)
	LoggerRequestInit(cfg)
	Logger.Info("Log-Instances Created!")
}

func LoggerDefaultInit(cfg zap.Config) {
	fileMethod := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/develop.log",
		MaxSize:    100, // megabytes
		MaxAge:     15,  // days
		MaxBackups: 3,
		LocalTime:  true,
		Compress:   false,
	})

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(cfg.EncoderConfig), os.Stdout, cfg.Level),
		zapcore.NewCore(zapcore.NewJSONEncoder(cfg.EncoderConfig), fileMethod, cfg.Level),
	)
	logInstance := zap.New(core, zap.Development())
	Logger = logInstance.Sugar()
}

func LoggerRequestInit(cfg zap.Config) {
	fileRequestMethod := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "logs/requests.log",
		MaxSize:    3,  // megabytes
		MaxAge:     15, // days
		MaxBackups: 3,
		LocalTime:  false,
		Compress:   true,
	})

	cores := []zapcore.Core{}
	if configProvider.GetConfig().GetBool("log.request.console") == true {
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(cfg.EncoderConfig), os.Stdout, cfg.Level))
	}
	if configProvider.GetConfig().GetBool("log.request.file") == true {
		cores = append(cores, zapcore.NewCore(zapcore.NewConsoleEncoder(cfg.EncoderConfig), fileRequestMethod, cfg.Level))
	}
	coreRequest := zapcore.NewTee(cores...)
	logRequestInstance := zap.New(coreRequest, zap.Development())
	LoggerRequest = logRequestInstance.Sugar()
}
