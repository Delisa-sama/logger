package logger

import (
	"database/sql"
	stash "github.com/Delisa-sama/logger/stash"
	"github.com/fatih/color"
	"io"
	"os"
	"time"
)

// Энумератор для уровней логирования
type LogLevel int

const (
	FATAL = iota // Фатальный уровень, вызывает падение приложения
	ERROR        // Ошибка
	WARN         // Предупреждение
	DEBUG        // Дебаг
	INFO         // Информационные логи
)

func (l LogLevel) String() string {
	return [...]string{"FATAL", "ERROR", "WARN", "DEBUG", "INFO", "STASH"}[l]
}

type Options struct {
	output         *os.File
	level          LogLevel
	colorize       bool
	colorFunctions WriterFunctions
	timeFormat     string
	stash          *stash.Stash
}

type Option func(o *Options)

type WriterFunctions map[LogLevel]func(w io.Writer, format string, v ...interface{})

// Default option values
var (
	DefaultWriterFunctions = WriterFunctions{
		FATAL: color.New(color.FgBlack, color.BgRed).FprintfFunc(),
		ERROR: color.New(color.FgRed).FprintfFunc(),
		WARN:  color.New(color.FgYellow).FprintfFunc(),
		DEBUG: color.New(color.FgCyan).FprintfFunc(),
		INFO:  color.New(color.FgWhite).FprintfFunc(),
	}
	DefaultOutput              = os.Stdout
	DefaultLevel      LogLevel = WARN
	DefaultTimeFormat          = time.RFC3339
)

// private logger options constructor
func newOptions(opts ...Option) Options {
	opt := Options{
		output:         DefaultOutput,
		level:          DefaultLevel,
		colorize:       false,
		colorFunctions: DefaultWriterFunctions,
		timeFormat:     DefaultTimeFormat,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// output option setter
func Output(file *os.File) Option {
	return func(o *Options) {
		o.output = file
	}
}

func Level(level LogLevel) Option {
	return func(o *Options) {
		o.level = level
	}
}

func Colorize(colorize bool) Option {
	return func(o *Options) {
		o.colorize = colorize
	}
}

func WriterFunc(writerFunc *WriterFunctions) Option {
	return func(o *Options) {
		if writerFunc == nil {
			return
		}
		for level, function := range *writerFunc {
			o.colorFunctions[level] = function
		}
	}
}

func StashOutput(db *sql.DB) Option {
	return func(o *Options) {
		o.stash = stash.NewStash(db)
	}
}
