package qodercnhooks

import (
	"encoding/json"
	"fmt"

	"github.com/hellolib/agent-notify/internal/notify"
)

// payload describes Qoder CN hooks event JSON delivered through stdin.
type payload struct {
	HookEventName string         `json:"hook_event_name"`
	SessionID     string         `json:"session_id"`
	CWD           string         `json:"cwd"`
	ToolName      string         `json:"tool_name"`
	ToolInput     map[string]any `json:"tool_input"`
}

func ParseMessage(data []byte) (notify.Message, error) {
	var p payload
	if err := json.Unmarshal(data, &p); err != nil {
		return notify.Message{}, err
	}

	switch p.HookEventName {
	case "PreToolUse":
		return notify.Message{
			Agent:     "qodercn",
			Event:     "permission_required",
			SessionID: p.SessionID,
			Workspace: p.CWD,
			Title:     notify.FormatTitle("qodercn", "permission_required"),
			Body:      fmt.Sprintf("工具: %s\n操作需要您的授权许可", fallbackToolName(p.ToolName)),
		}, nil
	case "Stop":
		return notify.Message{
			Agent:     "qodercn",
			Event:     "run_completed",
			SessionID: p.SessionID,
			Workspace: p.CWD,
			Title:     notify.FormatTitle("qodercn", "run_completed"),
			Body:      notify.DefaultBody("run_completed"),
		}, nil
	case "PostToolUseFailure":
		return notify.Message{
			Agent:     "qodercn",
			Event:     "run_failed",
			SessionID: p.SessionID,
			Workspace: p.CWD,
			Title:     notify.FormatTitle("qodercn", "run_failed"),
			Body:      fmt.Sprintf("工具: %s\n错误: 操作失败", fallbackToolName(p.ToolName)),
		}, nil
	default:
		return notify.Message{}, fmt.Errorf("unsupported hook event: %s", p.HookEventName)
	}
}

func fallbackToolName(name string) string {
	if name == "" {
		return "未知工具"
	}
	return name
}
