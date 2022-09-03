package utilz

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

func If[T any](condition bool, ifOutput T, elseOutput T) T {
	if condition {
		return ifOutput
	}

	return elseOutput
}

func Exec(command string, args []string, stdin []byte) ([]byte, error) {
	path, err := exec.LookPath(command)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(path, args...)
	if len(stdin) > 0 {
		var output bytes.Buffer
		cmd.Stdin = bytes.NewReader(stdin)
		cmd.Stdout = &output
		cmd.Stderr = &output
		err := cmd.Run()
		return output.Bytes(), err
	} else {
		return cmd.CombinedOutput()
	}
}

func GetEnvString(prefix, name string) string {
	return os.Getenv(strings.ToUpper(prefix + "_" + name))
}

func GetEnvBool(prefix, name string) bool {
	return strings.EqualFold(GetEnvString(prefix, name), "true")
}
