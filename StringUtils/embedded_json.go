package StringUtils

import (
	"strings"
)

func EmbeddedJson(jsonStr string) string {
	t := strings.Replace(jsonStr, "\"", "\\\"", -1)

	return trings.Replace(t, "\\\\", "\\" -1)
}
