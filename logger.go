package logger

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"time"
)

type LogWriter interface {
	// Логи уровня ERROR
	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	// Логи уровня DEBUG
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})

	// Логи уровня FATAL
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})

	// Логи уровня INFO
	Info(v ...interface{})
	Infof(format string, v ...interface{})

	// Логи уровня WARN
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
}

// escape последовательность для очищения значения текущего цвета
const escape = "\x1b"

type Logger struct {
	opts Options
}

// инстанс синглтона
var logger *Logger

// публичный конструктор
func NewLogger(opts ...Option) *Logger {
	options := newOptions(opts...)
	return &Logger{opts: options}
}

// Конструктор синглтона
func Init(opts ...Option) {
	if logger == nil {
		logger = NewLogger(opts...)
	}
}

// геттер для синглтона
func GetLogger() *Logger {
	return logger
}

// сброс цвета
func (l *Logger) resetColor() {
	_, _ = fmt.Fprintf(l.opts.output, "%s[%dm", escape, color.Reset)
}

// напечатать сообщение в output логера
// формат сообщения: [TIME][LEVEL] MESSAGE
// принимает на вход формат в стиле Printf для формирования MESSAGE
func (l *Logger) printf(level LogLevel, format string, v ...interface{}) {
	format = fmt.Sprintf("[%s][%s] %s", time.Now().Format(l.opts.timeFormat), level, format) // TODO: make format abstraction
	if l.opts.colorize {
		l.opts.colorFunctions[level](l.opts.output, format, v...)
		l.resetColor()
	} else {
		_, _ = fmt.Fprintf(l.opts.output, format, v...)
	}
	fmt.Println()
}

// напечатать сообщение в stderr
// формат сообщения: [TIME][LEVEL] MESSAGE
// принимает на вход формат в стиле Printf для формирования MESSAGE
func (l *Logger) errorf(level LogLevel, format string, v ...interface{}) {
	format = fmt.Sprintf("[%s][%s] %s", time.Now().Format(l.opts.timeFormat), level, format) // TODO: make format abstraction
	_ = fmt.Errorf(format, v...)
}

// напечатать ошибку
// печатает ошибку в stderr и в output логера
func (l *Logger) Error(v ...interface{}) {
	l.Errorf("%s", v...)
}

// напечатать ошибку с форматированием
// печатает ошибку в stderr и в output логера
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.opts.level >= ERROR {
		l.errorf(ERROR, format, v...)
		l.printf(ERROR, format, v...)
	}
}

// напечатать дебаг в output логера
func (l *Logger) Debug(v ...interface{}) {
	l.Debugf("%s", v...)
}

// напечатать дебаг с форматированием в output логера
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.opts.level >= DEBUG {
		l.printf(DEBUG, format, v...)
	}
}

// напечатать фатальную ошибку и завершить работу программы
// печатает сообщение в stderr и в output логера
// затем завершает программу с кодом возврата 1 `os.Exit(1)`
func (l *Logger) Fatal(v ...interface{}) {
	l.Fatalf("%s", v...)
}

// напечатать фатальную ошибку с форматированием и завершить работу программы
// печатает сообщение в stderr и в output логера
// затем завершает программу с кодом возврата 1 `os.Exit(1)`
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if l.opts.level >= FATAL {
		l.errorf(FATAL, format, v...)
		l.printf(FATAL, format, v...)
		os.Exit(1)
	}
}

// напечатать информационное сообщение в output логера
func (l *Logger) Info(v ...interface{}) {
	l.Infof("%s", v...)
}

// напечатать информационное сообщение с форматированием в output логера
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.opts.level >= INFO {
		l.printf(INFO, format, v...)
	}
}

// напечатать предупреждение в output логера
func (l *Logger) Warn(v ...interface{}) {
	l.Warnf("%s", v...)
}

// напечатать предупреждение с форматированием в output логера
func (l *Logger) Warnf(format string, v ...interface{}) {
	if l.opts.level >= WARN {
		l.printf(WARN, format, v...)
	}
}
