package utilz

import (
	"os/exec"
)

func If[T any](condition bool, ifOutput T, elseOutput T) T {
	if condition {
		return ifOutput
	}

	return elseOutput
}

func Exec(command string, args ...string) (string, error) {
	path, err := exec.LookPath(command)
	if err != nil {
		return "", nil
	}
	output, err := exec.Command(path, args...).CombinedOutput()
	if err != nil {
		return "", nil
	}
	return string(output), nil
}
