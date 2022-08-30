package utilz

import (
	"bufio"
	"strings"
)

type KeyVals map[string]string

func (m KeyVals) Get(key, defaultValue string) string {
	if val, ok := m[key]; ok {
		return val
	} else {
		return defaultValue
	}
}

func (m KeyVals) GetString(key string) string {
	return m.Get(key, "")
}

func (m KeyVals) GetBool(key string) bool {
	return strings.EqualFold(m.GetString(key), "true")
}

func WriteKeyValueToString(data KeyVals) string {
	var sb strings.Builder
	for k, v := range data {
		sb.WriteString(k)
		sb.WriteRune('=')
		sb.WriteString(v)
		sb.WriteRune('\n')
	}
	return sb.String()
}

func ReadKeyValueFromString(data string) (KeyVals, error) {
	ret := make(KeyVals)
	bs := bufio.NewScanner(strings.NewReader(data))
	bs.Split(bufio.ScanLines)
	for bs.Scan() {
		line := strings.TrimSpace(bs.Text())
		if len(line) == 0 { // skip empty line
			continue
		}
		if strings.HasPrefix(line, "#") { // ignore comment
			continue
		}
		markPos := strings.IndexRune(line, '=')
		if markPos == -1 {
			ret[line] = "true" // bool
		} else {
			key := line[:markPos]
			val := line[markPos+1:]
			ret[key] = val
		}
	}
	return ret, nil
}
