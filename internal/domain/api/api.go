package api

import "context"

// these functions print the result
type Api interface {
	ListTasks(ctx context.Context) error
	PushTask(ctx context.Context, title string, priority float64) error
	PopTask(ctx context.Context, priotirty float64) error
	// CleanTasks(ctx context.Context) error
}
