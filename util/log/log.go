package jasonlog

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *zap.SugaredLogger

func init() {
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel
	})
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.ErrorLevel
	})

	panicLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.PanicLevel
	})
	core := zapcore.NewTee(
		getZapCore(getWarnFilePath(), warnLevel),
		getZapCore(getInfoFilePath(), infoLevel),
		getZapCore(getErrorFilePath(), errorLevel),
		getZapCore(getPanicFilePath(), panicLevel),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	defer logger.Sync()
	log = logger.Sugar()
}

func getZapCore(path string, level zapcore.LevelEnabler) zapcore.Core {
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
	format := &lumberjack.Logger{
		Filename:  path,
		MaxSize:   1024, //MB
		LocalTime: true,
		Compress:  true,
	}
	w := zapcore.AddSync(format)
	//config := zap.NewProductionEncoderConfig()
	//config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		encoder,
		w,
		level,
	)
	return core
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Info(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func getInfoFilePath() string {

	logfile := getCurrentDirectory() + "/../logs/" + "info.log." + time.Now().Format("2006-01-02")
	return logfile
}

func getErrorFilePath() string {
	logfile := getCurrentDirectory() + "/../logs/"+ "error.log." + time.Now().Format("2006-01-02")
	return logfile
}

func getWarnFilePath() string {
	logfile := getCurrentDirectory() + "/../logs/" + "warn.log." + time.Now().Format("2006-01-02")
	return logfile
}

func getPanicFilePath() string {
	logfile := getCurrentDirectory() + "/../logs/" + "panic.log." + time.Now().Format("2006-01-02")
	return logfile
}

func getAppname() string {
	full := os.Args[0]
	full = strings.Replace(full, "\\", "/", -1)
	splits := strings.Split(full, "/")
	if len(splits) >= 1 {
		name := splits[len(splits)-1]
		name = strings.TrimSuffix(name, ".exe")
		return name
	}

	return ""
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	log.Warnf(template, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	log.Panicf(template, args...)
}
