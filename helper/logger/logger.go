/*
 * @Author: Bin
 * @Date: 2023-04-09
 * @FilePath: /gpt-zmide-server/helper/logger/logger.go
 */
package logger

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var LOG_FILE_PATH = "./debug.log"

var (
	_logger *zap.Logger
	once    sync.Once
)

var (
	Debug func(msg string, fields ...zap.Field)
	Info  func(msg string, fields ...zap.Field)
	Warn  func(msg string, fields ...zap.Field)
	Error func(msg string, fields ...zap.Field)
	Fatal func(msg string, fields ...zap.Field)
)

func InitLogger() {
	once.Do(func() {
		_logger = loadLogConfig()
	})

	Debug = _logger.Debug
	Info = _logger.Info
	Warn = _logger.Warn
	Error = _logger.Error
	Fatal = _logger.Fatal
}

func loadLogConfig() *zap.Logger {
	encoder := getEncoder()
	sync := getWriteSync()
	core := zapcore.NewCore(encoder, sync, zapcore.DebugLevel)

	if os.Getenv("GIN_DEBUG") == "true" {
		// 开启开发模式，堆栈跟踪
		caller := zap.AddCaller()        // 开启文件及行号
		development := zap.Development() // 构造日志对象
		return zap.New(core, caller, development)
	} else {
		return zap.New(core)
	}
}

func getEncoder() zapcore.Encoder {
	stdEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	return zapcore.NewConsoleEncoder(stdEncoderConfig)
}

func getWriteSync() zapcore.WriteSyncer {
	loggerFileWriter := lumberjack.Logger{
		Filename:   LOG_FILE_PATH, // 日志文件路径
		MaxSize:    50,            // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 1,             // 日志文件最多保存多少个备份
		MaxAge:     30,            // 文件最多保存多少天
		Compress:   true,          // 是否压缩
	}
	syncConsole := zapcore.AddSync(os.Stdout)
	syncFile := zapcore.AddSync(&loggerFileWriter)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
