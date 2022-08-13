package utilz

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
)

func ToJSON(data any) ([]byte, error) {
	return jsoniter.Marshal(data)
}

func FromJSON(data []byte, v any) error {
	return jsoniter.Unmarshal(data, v)
}

func ToJSONString(data any) (string, error) {
	return jsoniter.MarshalToString(data)
}

func FromJSONString(data string, v any) error {
	return jsoniter.UnmarshalFromString(data, v)
}

func ToJSONStringNoError(data any) string {
	result, _ := jsoniter.MarshalToString(data)
	return result
}

func init() {
	extra.RegisterFuzzyDecoders()
}
