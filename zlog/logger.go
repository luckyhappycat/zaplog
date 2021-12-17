package zlog

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.Logger        = CreateLogger("")
	Sugar  *zap.SugaredLogger = Logger.Sugar()

	//configMap = &sync.Map{}
)

//Debug zap模式打印日志
func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

//Debugf printf模式打印日志
func Debugf(template string, args ...interface{}) {
	Sugar.Debugf(template, args...)
}

//Info zap模式打印日志
func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

//Infof printf模式打印日志
func Infof(template string, args ...interface{}) {
	Sugar.Infof(template, args...)
}

//Warn zap模式打印日志
func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

//Warnf printf模式打印日志
func Warnf(template string, args ...interface{}) {
	Sugar.Warnf(template, args...)
}

//Error zap模式打印日志
func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

//Errorf printf模式打印日志
func Errorf(template string, args ...interface{}) {
	Sugar.Errorf(template, args...)
}

//Fatal zap模式打印日志
func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}

//Fatalf printf模式打印日志
func Fatalf(template string, args ...interface{}) {
	Sugar.Fatalf(template, args...)
}

//Panic zap模式打印日志
func Panic(msg string, fields ...zap.Field) {
	Logger.Panic(msg, fields...)
}

//Panicf printf模式打印日志
func Panicf(template string, args ...interface{}) {
	Sugar.Panicf(template, args...)
}

//FatalWithError 便于快速报错退出的函数
func FatalWithError(err error) {
	Fatal("ERROR!", zap.Error(err))
}

func CreateLogger(name string) *zap.Logger {
	enc := os.Getenv("LOG_ENCODER")
	if enc == "" {
		enc = "console"
	}

	cfg := zap.NewProductionConfig()
	cfg.Encoding = enc
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.ConsoleSeparator = " "
	if enc == "console" {
		cfg.EncoderConfig.EncodeName = func(s string, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString("[" + s + "]")
		}
		cfg.EncoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString("(" + caller.TrimmedPath() + ")")
		}
	}
	cfg.DisableStacktrace = true
	cfg.Level.SetLevel(getLoggerLevelByName(name))

	l, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	return l.Named(name)
}

func getLoggerLevelByName(name string) zapcore.Level {
	s := os.Getenv("LOG_LEVEL")
	mods := strings.Split(s, ",")
	lvMap := map[string]zapcore.Level{}
	for _, m := range mods {
		t := strings.SplitN(strings.TrimSpace(m), "=", 2)
		var modName string
		var lv string
		if len(t) == 1 {
			modName = ""
			lv = t[0]
		} else {
			modName = t[0]
			lv = t[1]
		}

		l := zap.InfoLevel
		err := l.UnmarshalText([]byte(lv))
		if err != nil {
			// pass
		}

		lvMap[modName] = l
	}

	if _, ok := lvMap[""]; !ok {
		lvMap[""] = zap.InfoLevel
	}

	l, ok := lvMap[name]
	if ok {
		return l
	} else {
		return lvMap[""]
	}
}
