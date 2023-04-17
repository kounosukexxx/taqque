package main

import (
	"fmt"
	"os"
	"strconv"

	taqqueCli "github.com/kounosukexxx/taqque/internal/services/cli"
	"github.com/urfave/cli/v2"
)

var listCmd = &cli.Command{
	Name:  "list",
	Usage: "list tasks",
	Action: func(ctx *cli.Context) error {
		cli, err := taqqueCli.NewCli()
		if err != nil {
			return err
		}

		err = cli.Api.ListTasks(ctx.Context)
		if err != nil {
			return err
		}

		return nil
	},
}

var pushCmd = &cli.Command{
	Name:  "push",
	Usage: "push a task with priority. by default, priority is 1",
	Action: func(ctx *cli.Context) error {
		taskTitle := ctx.Args().Get(0)
		priorityStr := ctx.Args().Get(1)

		priority := 1.0
		switch {
		case priorityStr == "":
			break
		default:
			var err error
			priority, err = strconv.ParseFloat(priorityStr, 64)
			if err != nil || priority < 0 {
				return fmt.Errorf("priority must be positive float number. got %s", priorityStr)
			}
		}

		cli, err := taqqueCli.NewCli()
		if err != nil {
			return err
		}

		err = cli.Api.PushTask(ctx.Context, taskTitle, priority)
		if err != nil {
			return err
		}

		return nil
	},
}

var popCmd = &cli.Command{
	Name:  "pop",
	Usage: "pop a task of specific priority",
	Action: func(ctx *cli.Context) error {
		priorityStr := ctx.Args().Get(0)

		priority := 1.0
		switch {
		case priorityStr == "":
			break
		default:
			var err error
			priority, err = strconv.ParseFloat(priorityStr, 64)
			if err != nil || priority < 0 {
				return fmt.Errorf("priority must be positive float number. got %s", priorityStr)
			}
		}
		cli, err := taqqueCli.NewCli()
		if err != nil {
			return err
		}

		err = cli.Api.PopTask(ctx.Context, priority)
		if err != nil {
			return err
		}

		return nil
	},
}

// var cleanCmd = &cli.Command{
// 	Name:  "clean",
// 	Usage: "clean all tasks",
// 	Action: func(ctx *cli.Context) error {
// 		cli := taqqueCli.NewCli()

// 		if err := cli.Api.CleanTasks(); err != nil {
// 			_, _ = fmt.Fprintf(os.Stderr, "%+v\n", err.Error())
// 			os.Exit(2)
// 		}

// 		return nil
// 	},
// }

func main() {
	app := cli.NewApp()
	app.Name = "taqque"
	app.Usage = listCmd.Usage
	app.Description = "This is a task manegement tool using queues with priority concept."
	app.DefaultCommand = "list"
	app.Commands = []*cli.Command{
		listCmd,
		pushCmd,
		popCmd,
		// cleanCmd,
	}
	app.SkipFlagParsing = true

	if err := app.Run(os.Args); err != nil {
		fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
}
