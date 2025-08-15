package pkg

import (
	"errors"
	"os"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	instance *zap.Logger
	once     sync.Once
)

// Config 日志配置
type Config struct {
	LogDir       string // 日志根目录
	BusinessLog  string // 业务日志文件名
	ErrorLog     string // 错误日志文件名
	MaxSize      int    // 日志文件最大大小(MB)
	MaxBackups   int    // 最大保留日志文件数
	MaxAge       int    // 日志保留天数
	Compress     bool   // 是否压缩日志
	Development  bool   // 是否开发模式
	JsonEncoding bool   // 是否使用JSON格式
	ShowCaller   bool   // 是否显示调用者信息
	LogLevel     string // 日志级别(debug/info/warn/error/dpanic/panic/fatal)
}

// InitLogger 初始化日志（单例模式）
func InitLogger(cfg Config) error {
	var initErr error
	once.Do(func() {
		// 设置默认值
		if cfg.LogDir == "" {
			cfg.LogDir = "./logs"
		}
		if cfg.BusinessLog == "" {
			cfg.BusinessLog = "business.log"
		}
		if cfg.ErrorLog == "" {
			cfg.ErrorLog = "error.log"
		}
		if cfg.MaxSize <= 0 {
			cfg.MaxSize = 100 // 默认100MB
		}
		if cfg.MaxBackups <= 0 {
			cfg.MaxBackups = 30 // 默认保留30个文件
		}
		if cfg.MaxAge <= 0 {
			cfg.MaxAge = 7 // 默认保留7天
		}

		// 创建日志目录
		businessDir := filepath.Join(cfg.LogDir, "business")
		errorDir := filepath.Join(cfg.LogDir, "error")

		if err := os.MkdirAll(businessDir, 0755); err != nil {
			initErr = errors.Join(errors.New("创建业务日志目录失败"), err)
			return
		}
		if err := os.MkdirAll(errorDir, 0755); err != nil {
			initErr = errors.Join(errors.New("创建错误日志目录失败"), err)
			return
		}

		// 设置日志级别
		level := zapcore.InfoLevel
		if err := level.Set(cfg.LogLevel); err != nil && cfg.LogLevel != "" {
			initErr = errors.Join(errors.New("设置日志级别失败"), err)
			return
		}

		// 开发模式直接输出到控制台
		if cfg.Development {
			instance = initDevLogger(cfg)
			return
		}

		// 生产环境日志配置
		instance = initProdLogger(cfg, businessDir, errorDir)
	})

	return initErr
}

func initDevLogger(cfg Config) *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)

	options := []zap.Option{
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	}

	return zap.New(core, options...)
}

func initProdLogger(cfg Config, businessDir, errorDir string) *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var encoder zapcore.Encoder
	if cfg.JsonEncoding {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 业务日志Writer (Info/Warn级别)
	businessWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(businessDir, cfg.BusinessLog),
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	})

	// 错误日志Writer (Error/Fatal级别)
	errorWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(errorDir, cfg.ErrorLog),
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	})

	// 业务日志Core (Info/Warn)
	businessCore := zapcore.NewCore(
		encoder,
		businessWriter,
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.InfoLevel && lvl <= zapcore.WarnLevel
		}),
	)

	// 错误日志Core (Error/Fatal)
	errorCore := zapcore.NewCore(
		encoder,
		errorWriter,
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		}),
	)

	// 合并Core
	core := zapcore.NewTee(businessCore, errorCore)

	// 添加选项
	options := []zap.Option{}
	if cfg.ShowCaller {
		options = append(options, zap.AddCaller())
	}
	options = append(options, zap.AddStacktrace(zapcore.ErrorLevel))

	return zap.New(core, options...)
}

// 确保日志已初始化（未初始化则用默认配置）
func ensureInitialized() {
	if instance == nil {
		_ = InitLogger(Config{})
	}
}

// 日志方法封装
func Debug(msg string, fields ...zap.Field) {
	ensureInitialized()
	instance.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	ensureInitialized()
	instance.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	ensureInitialized()
	instance.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	ensureInitialized()
	instance.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	ensureInitialized()
	instance.Fatal(msg, fields...)
}

// Sync 刷新日志缓冲区
func Sync() error {
	if instance == nil {
		return nil
	}
	return instance.Sync()
}
