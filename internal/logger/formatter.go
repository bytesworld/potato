package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

const (
	defaultLogFormat = ""[%levelname%] %time% P_%(process) - d_T_%(thread)d ' \
'[ %(pathname)s:%(lineno)d ]: %(message)s'
	defaultTimestampFormat = "2006-01-02 15:04:05"
)
type PotatoFormatter struct {
	TimestampFormat string
	LogFormat string
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

	output = strings.Replace(output, "%msg%", entry.Message, 1)

	level := strings.ToUpper(entry.Level.String())
	output = strings.Replace(output, "%lvl%", level, 1)

	for k, val := range entry.Data {
		switch v := val.(type) {
		case string:
			output = strings.Replace(output, "%"+k+"%", v, 1)
		case int:
			s := strconv.Itoa(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		case bool:
			s := strconv.FormatBool(v)
			output = strings.Replace(output, "%"+k+"%", s, 1)
		}
	}

	return []byte(output), nil
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	level := levelList[int(entry.Level)]
	strList := strings.Split(entry.Caller.File, "/")
	fileName := strList[len(strList)-1]
	b.WriteString(fmt.Sprintf("%s - %s - [line:%d] - %s - %s\n",
		entry.Time.Format("2006-01-02 15:04:05,678"), fileName,
		entry.Caller.Line, level, entry.Message))
	return b.Bytes(), nil
}


