package log

import (
    rotatelogs "github.com/lestrrat-go/file-rotatelogs"
    "github.com/sirupsen/logrus"
    "os"
    "fmt"
)

type LogrusLogger struct {
    *logrus.Logger
}

func NewLogrusLogger() *LogrusLogger {
    Logger := &LogrusLogger{
        Logger: logrus.New(),
    }

    Logger.Out = os.Stdout
    Logger.SetLevel(logrus.TraceLevel)
    Logger.Formatter = &logrus.TextFormatter{
        DisableColors:   true,
        TimestampFormat: "2006/01/02 15:04:05"}
    Logger.Println("logger init success")
    return Logger
}

func (l *LogrusLogger) SetWriteToFile(path string, filename string) {
    _, err := os.Stat(path)
    if err != nil {
        os.MkdirAll(path, os.ModePerm)
    }
    full := fmt.Sprintf("%s/%s.%Y%m.log", path, filename)
    ll, _ := rotatelogs.New(full, rotatelogs.WithRotationCount(100))
    l.Out = ll
}

func (l *LogrusLogger) Log(v ...interface{}) {
    if l.Logger != nil {
        l.Print(v...)
    } else {
        logrus.Print(v...)
    }
}

func (l *LogrusLogger) Output(n int, s string) error {
    if l.Logger != nil {
        l.Print(n, s)
    } else {
        logrus.Print(n, s)
    }
    return nil
}

func (l *LogrusLogger) Logf(format string, v ...interface{}) {
    if l.Logger != nil {
        l.Printf(format, v...)
    } else {
        logrus.Printf(format, v...)
    }
}
