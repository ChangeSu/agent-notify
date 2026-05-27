package qodercnhooks

import (
	"context"
	"fmt"
	"io"

	"github.com/hellolib/agent-notify/internal/agenthooks"
	"github.com/hellolib/agent-notify/internal/config"
	"github.com/hellolib/agent-notify/internal/state"
)

func Handle(ctx context.Context, cfg config.Config, statePath, logPath string, stdin io.Reader) error {
	data, err := io.ReadAll(stdin)
	if err != nil {
		return state.AppendLog(logPath, fmt.Sprintf("read stdin error: %v", err))
	}

	msg, err := ParseMessage(data)
	if err != nil {
		return state.AppendLog(logPath, fmt.Sprintf("skip event: %v", err))
	}

	return agenthooks.Dispatch(ctx, cfg, statePath, logPath, msg)
}
