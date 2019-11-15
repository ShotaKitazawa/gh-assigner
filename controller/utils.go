package controller

import (
	"strings"
)

func trimNewlineChar(str string) string {
	str = strings.Replace(str, "\n", " ", -1)
	str = strings.Replace(str, "\r", " ", -1)
	str = strings.Replace(str, "\r\n", " ", -1)
	return str
}
