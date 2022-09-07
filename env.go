package utilz

import (
	"os"
	"strings"
)

func LoadEnvFile(filename string) error {
	content, err := ReadFileToString(filename)
	if err != nil {
		return err
	}
	envs, err := ReadKeyValueFromString(content)
	if err != nil {
		return err
	}
	for k, v := range envs {
		err = os.Setenv(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetEnvString(name string) string {
	return os.Getenv(strings.ToUpper(name))
}

func GetEnvBool(name string) bool {
	return StringToBool(GetEnvString(name))
}
