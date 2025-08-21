package log

import (
	"log/slog"

	"github.com/areon546/go-files/files"
)

var (
	logFileName              = "./nova-skins.log"
	logFile                  = files.NewTextFile(logFileName)
	logger      *slog.Logger = newLogger(slog.LevelDebug)
)

func handlerOptions(lvl slog.Level) *slog.HandlerOptions {
	return &slog.HandlerOptions{AddSource: false, Level: lvl} // AddSource false because if true, it considers this log file as the source.
}

func newLogger(lvl slog.Level) *slog.Logger {
	return slog.New(slog.NewTextHandler(files.NewTextFile(logFileName), handlerOptions(lvl)))
}

func SetLogger(lvl slog.Level) {
	logger = newLogger(lvl)
}

func SetLogFileName(filename string) {
	logFileName = filename
}

func ClearLogFile() {
	logFile.ClearFile()
}

// INFO WARNING DEBUG ERROR

// Levels:
// DEBUG
// INFO
// WARNING
// ERROR

func Info(msg string, args ...any) {
	logger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	logger.Debug(msg, args...)
}

func Debug(msg string, args ...any) {
	logger.Debug(msg, args...)
}

func Error(msg string, args ...any) {
	logger.Debug(msg, args...)
}
