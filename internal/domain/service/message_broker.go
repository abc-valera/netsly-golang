package service

import (
	"context"
)

type Priority string

const (
	Low      Priority = "low"
	Default  Priority = "default"
	Critical Priority = "critical"
)

type MessageBroker interface {
	SendEmailTask(ctx context.Context, priority Priority, email Email) error
}
