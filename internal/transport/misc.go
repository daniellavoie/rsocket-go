package transport

import (
	"strings"
)

func IsClosedErr(err error) bool {
	if err == nil {
		return false
	}
	if strings.Contains(err.Error(), "use of closed network connection") {
		return true
	}
	return false
}