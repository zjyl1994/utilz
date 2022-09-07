package utilz

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func ByteSizeString(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

func EscapeSQLString(sql string) string {
	escapeMap := map[rune]rune{
		0:      '0',
		'\n':   'n',
		'\r':   'r',
		'\\':   '\\',
		'\'':   '\'',
		'"':    '"',
		'\032': 'Z',
	}

	var sb strings.Builder
	for _, r := range sql {
		if escape, ok := escapeMap[r]; ok {
			sb.WriteRune('\\')
			sb.WriteRune(escape)
		} else {
			sb.WriteRune(r)
		}
	}

	return sb.String()
}

func StringToBool(s string) bool {
	return strings.EqualFold(s, "true")
}

func Base64Encode(data []byte) string {
	return base64.RawStdEncoding.WithPadding('=').EncodeToString(data)
}

func Base64Decode(data string) ([]byte, error) {
	return base64.RawStdEncoding.WithPadding('=').DecodeString(data)
}
