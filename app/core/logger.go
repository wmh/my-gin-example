package core

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var commonLogger zerolog.Logger

func init() {
	logFile := ConfString("logs.common_log")

	// prepare directory
	dir, _ := filepath.Split(logFile)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	if ConfBool("logs.stdout_only") {
		commonLogger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	} else {
		f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err.Error())
		}
		commonLogger = zerolog.New(f).With().Timestamp().Logger()
	}
}

// CommonLog write common log with interface
func CommonLog(level zerolog.Level, tag string, msg interface{}, where interface{}) {
	now := time.Now()
	var event *zerolog.Event
	if level == zerolog.ErrorLevel {
		event = commonLogger.Error()
	} else {
		event = commonLogger.Log()
	}
	event.
		Str("date", now.Format("2006-01-02")).
		Str("datetime", now.Format("2006-01-02 15:04:05")).
		Str("tag", tag).
		Interface("message", msg)

	if where != nil {
		event.Interface("where", where)
	}
	event.Msg("")
}

// Log -
func Log(tag string, msg interface{}) {
	CommonLog(zerolog.NoLevel, tag, msg, nil)
}

// ErrorLog -
func ErrorLog(tag string, msg interface{}) {
	CommonLog(zerolog.ErrorLevel, tag, msg, whereAmI())
}

func whereAmI(depthList ...int) interface{} {
	var depth int
	if depthList == nil {
		depth = 3
	} else {
		depth = depthList[0]
	}
	function, file, line, _ := runtime.Caller(depth)
	out := dump{chopPath(file), runtime.FuncForPC(function).Name(), line}
	return out
}

type dump struct {
	File     string
	Function string
	Line     int
}

// return the source filename after the last slash
func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	}
	return original[i+1:]

}
