package qodercnhooks

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/hellolib/agent-notify/internal/common"
)

// BuildHookSettings generates the settings structure needed for Qoder CN settings.json.
// Qoder CN supports PreToolUse (permission_required), Stop (run_completed),
// and PostToolUseFailure (run_failed).
func BuildHookSettings(binaryPath string) map[string]any {
	binaryPath = common.ResolveBinaryPath(binaryPath)
	command := binaryPath + " handle-qodercn-hook"

	buildEntry := func() []map[string]any {
		return []map[string]any{
			{
				"hooks": []map[string]any{
					{
						"type":    "command",
						"command": command,
					},
				},
			},
		}
	}

	return map[string]any{
		"hooks": map[string]any{
			"PreToolUse":         buildEntry(),
			"Stop":               buildEntry(),
			"PostToolUseFailure": buildEntry(),
		},
	}
}

func Install(path string, binaryPath string) error {
	settings := map[string]any{}

	data, err := os.ReadFile(path)
	if err == nil {
		if err := json.Unmarshal(data, &settings); err != nil {
			return err
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	builtHooks := BuildHookSettings(binaryPath)["hooks"].(map[string]any)
	existingHooks, _ := settings["hooks"].(map[string]any)
	if existingHooks == nil {
		existingHooks = map[string]any{}
	}
	for key, value := range builtHooks {
		existingHooks[key] = value
	}
	settings["hooks"] = existingHooks

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	out, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, out, 0o644)
}
