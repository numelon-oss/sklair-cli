// TAKEN FROM https://github.com/numelon-bespoke/sunc-chan
// Copyright applies. This logger is proprietary software unless both signatories under the bespoke project agreement
// have agreed to make this code publicly available to other projects both within the Bespoke program,
// and to external projects unrelated to Numelon or the second signatory.

// 30/11/2025 Awaiting approval

package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Green  = "\033[32m"
	Cyan   = "\033[36m"
)

type LogLevel uint8

const (
	LevelNone LogLevel = iota
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
)

var levelTags = []struct {
	Raw    string
	Colour string
}{
	{"[NONE]", Reset},
	{"[ERROR]", Red},
	{"[WARN]", Yellow},
	{"[INFO]", Green},
	{"[DEBUG]", Cyan},
}

// Logger is a per-instance logger with level filtering and dual output
type Logger struct {
	level    LogLevel
	format   string
	file     *os.File
	stdout   io.Writer
	filePath string
}

// New Creates a new logger instance
func New(level LogLevel, dateTimeFormat string, filePath string) *Logger {
	var file *os.File
	var err error

	if filePath != "" {
		err = os.MkdirAll(filepath.Dir(filePath), 0755)
		if err != nil {
			fmt.Fprintf(os.Stderr, "logger: failed to create log directory: %v\n", err)
			os.Exit(1)
		}

		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "logger: failed to open log file: %v\n", err)
			os.Exit(1)
		}
	}

	return &Logger{
		level:    level,
		format:   dateTimeFormat,
		file:     file,
		stdout:   os.Stdout,
		filePath: filePath,
	}
}

func (l *Logger) Close() error {
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}

func (l *Logger) log(level LogLevel, format string, args ...any) {
	if l.level < level {
		return
	}

	tag := levelTags[level]
	timestamp := time.Now().Format(l.format)
	formatted := fmt.Sprintf(format, args...)

	// coloured stdout
	coloured := fmt.Sprintf("%s%s%s | %s", tag.Colour, tag.Raw, Reset, formatted)
	line := fmt.Sprintf("%s | %s\n", timestamp, coloured)
	fmt.Fprint(l.stdout, line)

	// plain file output
	if l.file != nil {
		rawLine := fmt.Sprintf("%s | %s | %s\n", timestamp, tag.Raw, formatted)
		l.file.WriteString(rawLine)
	}
}

// shortcut methods
func (l *Logger) Error(format string, args ...any)   { l.log(LevelError, format, args...) }
func (l *Logger) Warning(format string, args ...any) { l.log(LevelWarning, format, args...) }
func (l *Logger) Info(format string, args ...any)    { l.log(LevelInfo, format, args...) }
func (l *Logger) Debug(format string, args ...any)   { l.log(LevelDebug, format, args...) }
func (l *Logger) P(format string, args ...any)       { l.log(LevelNone, format, args...) }

// shared logger
var shared *Logger

func InitShared(level LogLevel, dateTimeFormat string, filePath string) {
	shared = New(level, dateTimeFormat, filePath)
}

// WILL LITERALLY EXPLODE IF SHARED NOT INITIALISED
func Error(format string, args ...any)   { shared.log(LevelError, format, args...) }
func Warning(format string, args ...any) { shared.log(LevelWarning, format, args...) }
func Info(format string, args ...any)    { shared.log(LevelInfo, format, args...) }
func Debug(format string, args ...any)   { shared.log(LevelDebug, format, args...) }
func P(format string, args ...any)       { shared.log(LevelNone, format, args...) }
func CloseShared() error                 { return shared.Close() }
