package agentintegrations

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hellolib/agent-notify/internal/common"
	"github.com/hellolib/agent-notify/internal/qodercnhooks"
)

// QoderCNIntegration implements Integration for Qoder CN (通义灵码).
type QoderCNIntegration struct{}

// NewQoderCNIntegration creates a new Qoder CN integration.
func NewQoderCNIntegration() *QoderCNIntegration {
	return &QoderCNIntegration{}
}

// Name returns the display name for Qoder CN.
func (q *QoderCNIntegration) Name() string {
	return "Qoder CN"
}

// DetectInstalled checks if Qoder CN is installed by looking for the ~/.lingma directory.
// Qoder CN does not have a standalone CLI, so we use the presence of the config directory
// as a heuristic for installation.
func (q *QoderCNIntegration) DetectInstalled() bool {
	home, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	_, err = os.Stat(filepath.Join(home, ".lingma"))
	return err == nil
}

// SettingsPath returns the path to Qoder CN's settings.json file.
func (q *QoderCNIntegration) SettingsPath(scope string) (string, error) {
	switch scope {
	case "user":
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, ".lingma", "settings.json"), nil
	case "project":
		return filepath.Join(".lingma", "settings.json"), nil
	default:
		return "", fmt.Errorf("unsupported scope: %s", scope)
	}
}

// Install configures Qoder CN to use agent-notify by setting up hooks.
func (q *QoderCNIntegration) Install(settingsPath, binaryPath string) error {
	return qodercnhooks.Install(settingsPath, common.ResolveBinaryPath(binaryPath))
}

// IsHookInstalled checks if agent-notify hooks are installed in the settings file.
func (q *QoderCNIntegration) IsHookInstalled(settingsPath string) (bool, error) {
	data, err := os.ReadFile(settingsPath)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	settings := map[string]any{}
	if err := json.Unmarshal(data, &settings); err != nil {
		return false, fmt.Errorf("failed to parse settings.json: %w", err)
	}

	hooks, ok := settings["hooks"].(map[string]any)
	if !ok {
		return false, nil
	}

	// Check if PreToolUse hook exists and contains handle-qodercn-hook
	pr, ok := hooks["PreToolUse"].([]any)
	if !ok || len(pr) == 0 {
		return false, nil
	}

	entry, ok := pr[0].(map[string]any)
	if !ok {
		return false, nil
	}

	hookList, ok := entry["hooks"].([]any)
	if !ok || len(hookList) == 0 {
		return false, nil
	}

	hook, ok := hookList[0].(map[string]any)
	if !ok {
		return false, nil
	}

	cmd, ok := hook["command"].(string)
	if !ok {
		return false, nil
	}

	return strings.Contains(cmd, "handle-qodercn-hook"), nil
}
