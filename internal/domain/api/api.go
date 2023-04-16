package api

import "context"

// these functions print the result
type Api interface {
	ListTasks(ctx context.Context) error
	PushTask(ctx context.Context, title string, priority int) error
	PopTask(ctx context.Context, priotirty int) error
	// CleanTasks(ctx context.Context) error
}
