package service

import (
	"context"
)

type TaskPriority string

const (
	Low      TaskPriority = "low"
	Default  TaskPriority = "default"
	Critical TaskPriority = "critical"
)

type ITaskQueuer interface {
	SendEmailTask(ctx context.Context, priority TaskPriority, email Email) error
}
