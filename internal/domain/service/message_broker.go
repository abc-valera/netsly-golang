package service

import (
	"context"
)

type Priority int

const (
	Low Priority = iota + 1
	Default
	Critical
)

type MessageBroker interface {
	SendEmailTask(ctx context.Context, priority Priority, email Email) error
}
