package common

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const DefaultBinaryPath = "agent-notify"

func ResolveBinaryPath(input string) string {
	input = strings.TrimSpace(input)
	if input != "" {
		return toUnixStylePath(input)
	}

	executablePath, err := os.Executable()
	if err == nil {
		if resolved, resolveErr := filepath.EvalSymlinks(executablePath); resolveErr == nil {
			return toUnixStylePath(resolved)
		}
		return toUnixStylePath(executablePath)
	}

	return DefaultBinaryPath
}

// toUnixStylePath normalizes Windows paths by converting backslashes to forward slashes.
// It preserves the drive letter format (e.g., C:\Users\... -> C:/Users/...) which is
// required by Windows-native tools such as Qoder CN hooks.
func toUnixStylePath(path string) string {
	if runtime.GOOS != "windows" {
		return path
	}

	// Convert backslashes to forward slashes while keeping the drive letter format.
	// This produces paths like "C:/Users/..." instead of "/c/Users/...".
	return strings.ReplaceAll(path, "\\", "/")
}
