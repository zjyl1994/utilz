package json

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

func Marshal(data any) ([]byte, error) {
	return jsoniter.Marshal(data)
}

func Unmarshal(data []byte, v any) error {
	return jsoniter.Unmarshal(data, v)
}

func MarshalToString(data any) (string, error) {
	return jsoniter.MarshalToString(data)
}

func UnmarshalToString(data string, v any) error {
	return jsoniter.UnmarshalFromString(data, v)
}

func MarshalToStringNoError(data any) string {
	result, _ := jsoniter.MarshalToString(data)
	return result
}

func init() {
	extra.RegisterFuzzyDecoders()
}
