package StringUtils

import (
	"strings"
)

func EmbeddedJson(jsonStr string) string {
	return strings.Replace(jsonStr, "\"", "\\\"", -1)
}
