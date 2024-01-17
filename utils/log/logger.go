package log

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

const (
	timeFormat = "2006-01-02 15:04:05"
	trace      = "traceId"
)

var (
	logger         zerolog.Logger
	filePrefixSkip int
)

func SetOutput(w io.Writer) zerolog.Logger {
	logger = zerolog.New(w).With().Timestamp().Logger()
	return logger
}

func SetFilePrefixSkip(val int) {
	filePrefixSkip = val
}

func Debug() *zerolog.Event {
	return logger.Debug().Str("time", Timer()).Str(trace, GetGoroutineId()).Str("caller", caller(filePrefixSkip))
}

func Info() *zerolog.Event {
	return logger.Info().Str("time", Timer()).Str(trace, GetGoroutineId()).Str("caller", caller(filePrefixSkip))
}

func Warn() *zerolog.Event {
	return logger.Warn().Str("time", Timer()).Str(trace, GetGoroutineId()).Str("caller", caller(filePrefixSkip))
}

func Error(err error) *zerolog.Event {
	return logger.Err(err).Str("time", Timer()).Str(trace, GetGoroutineId()).Str("caller", caller(filePrefixSkip))
}

func Fatal() *zerolog.Event {
	return logger.Fatal().Str("time", Timer()).Str(trace, GetGoroutineId()).Str("caller", caller(filePrefixSkip))
}

func GetGoroutineId() string {
	goroutineId := strconv.FormatUint(GetGID(), 10)
	return goroutineId
}

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func caller(prefixSkip int) string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	// fmt.Printf("file:%v line:%v\n", file, line)
	var b strings.Builder
	if prefixSkip > 0 {
		fileFields := strings.Split(file, "/")
		for idx, fileField := range fileFields {
			if idx < prefixSkip+1 {
				continue
			}
			if idx > prefixSkip+1 {
				b.WriteString("/")
			}
			b.WriteString(fileField)
		}
	} else {
		b.WriteString(file)
	}
	b.WriteString(fmt.Sprintf(":%v", line))
	return b.String()
}

func Timer() string {
	return time.Now().Format(timeFormat)
}

func init() {
	zerolog.CallerSkipFrameCount = 3
	zerolog.TimeFieldFormat = timeFormat

	logger = zerolog.New(os.Stdout).With().Logger()
}
