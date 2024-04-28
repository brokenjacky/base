package log

import (
    "io"
    "os"
    "path/filepath"
    "reflect"
    "github.com/natefinch/lumberjack"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "strings"
    "fmt"
    "github.com/go-kratos/kratos/v2/log"
)

const (
    cConsole = "console"
    cFile    = "file"
    cDebug   = "debug"
    cWarn    = "warn"
    cInfo    = "info"
    cError   = "error"

    cjson = "json"
)

type ZapLogger struct {
    zap.Logger
}

// LoggerConfig config for every logger
type LoggerConfig struct {
    Writer       string      `yaml:"writer"` // Writer console, file
    Level        string      `yaml:"level"`  // 日志输出级别
    Format       string      `yaml:"format"` // 日志输出格式, console, json
    WriterConfig WriteConfig `yaml:"writer_config"`
}

type WriteConfig struct {
    LogPath    string `yaml:"log_path"`    // 本地文件日志路径
    Filename   string `yaml:"filename"`    // 本地文件日志文件名
    RollType   string `yaml:"roll_type"`   // 文件滚动类型,size为按大小滚动
    MaxAge     int    `yaml:"max_age"`     // 最大日志保留天数
    MaxSize    int    `yaml:"max_size"`    // 本地文件滚动日志的大小 单位 MB
    MaxBackups int    `yaml:"max_backups"` // 最大日志文件数
}

func NewDefaultLoggerConfig(path string, filename string) *LoggerConfig {
    config := &LoggerConfig{
        Writer: cFile,
        Level:  cDebug,
        Format: cConsole,
        WriterConfig: WriteConfig{
            LogPath:    path,
            Filename:   fmt.Sprintf("%s.log", filename),
            RollType:   "size",
            MaxAge:     0,
            MaxSize:    100,
            MaxBackups: 100,
        },
    }

    if len(path) == 0 {
        config.Writer = cConsole
    }

    return config
}

func (cnf *LoggerConfig) getLogLever() zapcore.Level {
    var level zapcore.Level
    switch strings.ToLower(cnf.Level) {
    case cDebug:
        level = zapcore.DebugLevel

    case cInfo:
        level = zapcore.InfoLevel

    case cWarn:
        level = zapcore.WarnLevel

    case cError:
        level = zapcore.ErrorLevel

    default:
        level = zapcore.DebugLevel
    }
    return level
}

func NewZapLog(cnf *LoggerConfig, newWriter io.Writer) *ZapLogger {
    os.MkdirAll(filepath.Dir(cnf.WriterConfig.LogPath), 0777)
    var core zapcore.Core
    switch cnf.Writer {
    case cFile:
        encoder := newEncoder(cnf.Format)
        filename := filepath.Join(cnf.WriterConfig.LogPath, cnf.WriterConfig.Filename)
        w := &lumberjack.Logger{
            Filename:   filename,
            MaxSize:    cnf.WriterConfig.MaxSize,
            MaxBackups: cnf.WriterConfig.MaxBackups,
            MaxAge:     cnf.WriterConfig.MaxAge,
            Compress:   false,
        }
        ws := zapcore.AddSync(w)

        if newWriter != nil && !reflect.ValueOf(newWriter).IsNil() {
            newSync := zapcore.AddSync(newWriter)
            ws = zap.CombineWriteSyncers(zapcore.AddSync(w), newSync)
        }

        core = zapcore.NewCore(encoder, ws, cnf.getLogLever())

    default:
        encoder := newEncoder(cnf.Format)
        ws := zapcore.Lock(os.Stderr)
        core = zapcore.NewCore(encoder, ws, cnf.getLogLever())
    }

    var zlog = zap.New(core)
    ZapLogger := &ZapLogger{
        Logger: *zlog,
    }
    return ZapLogger
}
func newEncoder(format string) zapcore.Encoder {
    encoderCnf := zapcore.EncoderConfig{
        TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "name",
        CallerKey:      "line",
        MessageKey:     "msg",
        FunctionKey:    "func",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,
        EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000000000"),
        EncodeDuration: zapcore.NanosDurationEncoder,
        EncodeCaller:   zapcore.FullCallerEncoder,
        EncodeName:     zapcore.FullNameEncoder,
    }

    var encoder zapcore.Encoder
    switch format {
    case cConsole:
        encoder = zapcore.NewConsoleEncoder(encoderCnf)

    case cjson:
        encoder = zapcore.NewJSONEncoder(encoderCnf)

    default:
        encoder = zapcore.NewConsoleEncoder(encoderCnf)
    }

    return encoder
}

func (l *ZapLogger) Log(level log.Level, keyvals ...interface{}) error {
    if len(keyvals) == 0 || len(keyvals)%2 != 0 {
        l.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
        return nil
    }

    var data []zap.Field
    for i := 0; i < len(keyvals); i += 2 {
        data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
    }

    switch level {
    case log.LevelDebug:
        l.Debug("", data...)
    case log.LevelInfo:
        l.Info("", data...)
    case log.LevelWarn:
        l.Warn("", data...)
    case log.LevelError:
        l.Error("", data...)
    case log.LevelFatal:
        l.Fatal("", data...)
    }
    return nil
}
