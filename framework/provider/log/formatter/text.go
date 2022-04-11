package formatter

import (
	"bytes"
	"fmt"
	"time"
	"webman/framework/contract"
)

func TextFormatter(level contract.LogLevel, t time.Time, msg string, fields map[string]interface{}) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	separator := "\t"

	prefix := Prefix(level)
	// 输出日志级别
	buffer.WriteString(prefix)
	buffer.WriteString(separator)
	// 输出时间
	ts := t.Format(time.RFC3339)
	buffer.WriteString(ts)
	buffer.WriteString(separator)
	// 输出msg
	buffer.WriteString("\"")
	buffer.WriteString(msg)
	buffer.WriteString("\"")
	buffer.WriteString(separator)
	// 输出附加信息
	buffer.WriteString(fmt.Sprint(fields))
	return buffer.Bytes(), nil
}
