package api

import (
	"context"
	"os"
	"strconv"
	"time"

	a "github.com/kounosukexxx/taqque/internal/domain/api"
	"github.com/kounosukexxx/taqque/internal/domain/model"
	"github.com/kounosukexxx/taqque/internal/usecase"
	"github.com/olekukonko/tablewriter"
)

var _ a.Api = (*api)(nil)

type api struct {
	taskUsecase usecase.TaskUsecase
}

func NewApi(taskUsecase usecase.TaskUsecase) a.Api {
	return &api{
		taskUsecase: taskUsecase,
	}
}

func (c *api) ListTasks(ctx context.Context) error {
	tasks, err := c.taskUsecase.ListTasks(ctx)
	if err != nil {
		return err
	}
	c.printTasks(tasks)
	return nil
}

func (c *api) PushTask(ctx context.Context, title string, priority int) error {
	tasks, err := c.taskUsecase.PushTask(ctx, title, priority)
	if err != nil {
		return err
	}
	c.printTasks(tasks)
	return nil
}

func (c *api) PopTask(ctx context.Context, priority int) error {
	tasks, err := c.taskUsecase.PopTask(ctx, priority, time.Now())
	if err != nil {
		return err
	}
	c.printTasks(tasks)
	return nil
}

// func (c *api) CleanTasks() error {
// 	return nil // TODO: impletement
// }

func (c *api) printTasks(tasks []*model.Task) {
	data := make([][]string, len(tasks))
	for i, task := range tasks {
		data[i] = []string{strconv.Itoa(i), strconv.Itoa(task.Priority), task.Title}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Index", "Priority", "Title"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
