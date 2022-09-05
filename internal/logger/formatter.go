package logger

import (
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

const (
	defaultLogFormat       = "[%levelname%] %time%  [%filename%:%lineno%]: %message%\n"
	defaultTimestampFormat = "2006-01-02 15:04:05"
)

type PotatoFormatter struct {
	TimestampFormat string
	LogFormat       string
}

func (m PotatoFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	output := m.LogFormat
	if output == "" {
		output = defaultLogFormat
	}

	timestampFormat := m.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	output = strings.Replace(output, "%time%", entry.Time.Format(timestampFormat), 1)

	output = strings.Replace(output, "%message%", entry.Message, 1)

	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, "%levelname%", level, 1)

	output = strings.Replace(output, "%lineno%", strconv.Itoa(entry.Caller.Line), 1)
	output = strings.Replace(output, "%filename%", entry.Caller.File, 1)
	return []byte(output), nil
}
