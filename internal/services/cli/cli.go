package cli

import (
	"os"

	a "github.com/kounosukexxx/taqque/internal/domain/api"
	"github.com/kounosukexxx/taqque/internal/handler/api"
	"github.com/kounosukexxx/taqque/internal/handler/db/sqlite3"
	"github.com/kounosukexxx/taqque/internal/usecase"
)

type Cli struct {
	Api a.Api
}

func NewCli() (*Cli, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	taskRepository, err := sqlite3.NewTaskRepository(home)
	if err != nil {
		return nil, err
	}

	taskUsecase := usecase.NewTaskUsecase(taskRepository)
	api := api.NewApi(*taskUsecase)
	return &Cli{
		Api: api,
	}, nil
}
