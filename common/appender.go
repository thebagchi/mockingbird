package common

import "bytes"

const INDENT = "    "

func Append(values ...string) string {
	var buffer bytes.Buffer
	for _, value := range values {
		buffer.WriteString(value)
	}
	return buffer.String()
}
