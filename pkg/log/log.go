package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type Logger interface {
	Printf(format string, v ...interface{})
	Print(v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	ErrorLevel
	DisabledLevel
)

var (
	DebugLogger = &logger{DebugLevel}
	InfoLogger  = &logger{InfoLevel}
	ErrorLogger = &logger{ErrorLevel}
)

type globalState struct {
	currentLevel  Level
	defaultLogger Logger
}

var (
	mu    sync.RWMutex
	state = globalState{
		currentLevel:  InfoLevel,
		defaultLogger: newDefaultLogger(os.Stderr),
	}
)

func globals() globalState {
	mu.RLock()
	defer mu.RUnlock()
	return state
}

func newDefaultLogger(w io.Writer) Logger {
	return log.New(w, "", log.Ldate|log.Ltime)
}

type logger struct {
	level Level
}

func (l *logger) Print(v ...interface{}) {
	g := globals()

	if l.level < g.currentLevel {
		return // Don't log at lower levels.
	}
	if g.defaultLogger != nil {
		g.defaultLogger.Print(append([]interface{}{toPrefix(l.level)}, v...)...)
	}
}

func (l *logger) Printf(format string, v ...interface{}) {
	g := globals()

	if l.level < g.currentLevel {
		return // Don't log at lower levels.
	}
	if g.defaultLogger != nil {
		g.defaultLogger.Printf(toPrefix(l.level)+format, v...)
	}
}

func (l *logger) Fatal(v ...interface{}) {
	g := globals()

	if g.defaultLogger != nil {
		g.defaultLogger.Fatal(append([]interface{}{toPrefix(l.level)}, v...)...)
	} else {
		log.Fatal(append([]interface{}{toPrefix(l.level)}, v...)...)
	}
}

func (l *logger) Fatalf(format string, v ...interface{}) {
	g := globals()

	if g.defaultLogger != nil {
		g.defaultLogger.Fatalf("[FATAL] "+format, v...)
	} else {
		log.Fatalf("[FATAL] "+format, v...)
	}
}

func toPrefix(level Level) string {
	switch level {
	case InfoLevel:
		return "[INFO] "
	case DebugLevel:
		return "[DEBUG] "
	case ErrorLevel:
		return "[ERROR] "
	case DisabledLevel:
		return "[DISABLE] "
	}
	return "[UNKNOWN] "
}

func toString(level Level) string {
	switch level {
	case InfoLevel:
		return "info"
	case DebugLevel:
		return "debug"
	case ErrorLevel:
		return "error"
	case DisabledLevel:
		return "disabled"
	}
	return "unknown"
}

func toLevel(level string) (Level, error) {
	switch level {
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	case "error":
		return ErrorLevel, nil
	case "disabled":
		return DisabledLevel, nil
	}
	return DisabledLevel, fmt.Errorf("invalid log level %q", level)
}

func GetLevel() string {
	g := globals()

	return toString(g.currentLevel)
}

func SetLevel(level string) error {
	l, err := toLevel(level)
	if err != nil {
		return err
	}
	mu.Lock()
	state.currentLevel = l
	mu.Unlock()
	return nil
}

func Debug(v ...interface{}) {
	DebugLogger.Print(v...)
}

func Debugf(format string, v ...interface{}) {
	DebugLogger.Printf(format, v...)
}

func Info(v ...interface{}) {
	InfoLogger.Print(v...)
}

func Infof(format string, v ...interface{}) {
	InfoLogger.Printf(format, v...)
}

func Error(v ...interface{}) {
	ErrorLogger.Print(v...)
}

func Errorf(format string, v ...interface{}) {
	ErrorLogger.Printf(format, v...)
}

func Fatal(v ...interface{}) {
	InfoLogger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	InfoLogger.Fatalf(format, v...)
}
