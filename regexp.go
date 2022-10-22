package utilz

import (
	"regexp"
	"sync"
)

var regexpCache sync.Map

func CompileRegexp(exp string) *regexp.Regexp {
	iface, found := regexpCache.Load(exp)
	if found {
		return iface.(*regexp.Regexp)
	} else {
		compile := regexp.MustCompile(exp)
		regexpCache.Store(exp, compile)
		return compile
	}
}

func RegexpFindString(exp, data string) [][]string {
	complie := CompileRegexp(exp)
	matchs := complie.FindAllStringSubmatch(data, -1)
	if matchs == nil {
		return nil
	}
	result := make([][]string, 0, len(matchs))
	for _, v := range matchs {
		result = append(result, v[1:])
	}
	return result
}
