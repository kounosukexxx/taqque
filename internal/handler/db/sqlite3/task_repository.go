package sqlite3

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/kounosukexxx/taqque/internal/domain/model"
	"github.com/kounosukexxx/taqque/internal/domain/repositories"
	_ "github.com/mattn/go-sqlite3"
)

// TODO: なんかいい感じにする。生SQLを書きたく無い。また、Scanあたりの重複をなくしたい

var (
	_ repositories.TaskRepository = (*taskRepository)(nil)

	createTaskTableSQL = `
	CREATE TABLE IF NOT EXISTS tasks (id text, title text, priority float, sort_key text, created_at text, updated_at text, deleted_at text, PRIMARY KEY(id));
	`
)

const taskTableName = "tasks"

type taskRepository struct {
	sqlite3 *sqlite3
}

// TODO: consider db.Close() and migration
// userHomeDir: /Users/{username}
func NewTaskRepository(userHomeDir string) (repositories.TaskRepository, error) {
	if err := os.Mkdir(getTaqqueFolderPath(userHomeDir), 0755); err != nil {
		if !errors.Is(err, fs.ErrExist) {
			return nil, err
		}
	}
	db, err := sql.Open("sqlite3", getTaqqueDBFilePath(userHomeDir, taskTableName))
	if err != nil {
		return nil, err
	}

	sqlite3 := newSqlite3(userHomeDir, db)
	if err = sqlite3.initDB(createTaskTableSQL); err != nil {
		return nil, err
	}

	return &taskRepository{
		sqlite3: sqlite3,
	}, nil
}

// TODO: define taskEntity type
func (r *taskRepository) convertRowstoModel(rows *sql.Rows) ([]*model.Task, error) {
	var tasks []*model.Task
	defer rows.Close()
	for rows.Next() {
		var id, title, sort_key, created_time, updated_time string
		var priority float64
		if err := rows.Scan(&id, &title, &priority, &sort_key, &created_time, &updated_time); err != nil {
			return nil, err
		}
		sortKey, err := r.sqlite3.parseTime(sort_key)
		if err != nil {
			return nil, fmt.Errorf("failed to parse sort_key: %w", err)
		}
		createdTime, err := r.sqlite3.parseTime(created_time)
		if err != nil {
			return nil, fmt.Errorf("failed to parse create_time: %w", err)
		}
		updatedTime, err := r.sqlite3.parseTime(updated_time)
		if err != nil {
			return nil, fmt.Errorf("failed to parse update_time: %w", err)
		}
		tasks = append(tasks, &model.Task{
			ID:         id,
			Title:      title,
			Priority:   priority,
			SortKey:    sortKey,
			CreateTime: createdTime,
			UpdateTime: updatedTime,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *taskRepository) GetMultiOrderByPriorityDescAndSortKeyAsc(ctx context.Context, limit *int) ([]*model.Task, error) {
	var rows *sql.Rows
	var err error
	if limit == nil {
		sqlStmt := `
		SELECT id, title, priority, sort_key, created_at, updated_at FROM tasks WHERE deleted_at == "" ORDER BY priority DESC, sort_key ASC
		`
		rows, err = r.sqlite3.db.Query(sqlStmt)
	} else {
		sqlStmt := `
		SELECT id, title, priority, sort_key, created_at, updated_at FROM tasks WHERE deleted_at == "" ORDER BY priority DESC, sort_key ASC LIMIT ?
		`
		rows, err = r.sqlite3.db.Query(sqlStmt, *limit)
	}
	if err != nil {
		return nil, err
	}

	return r.convertRowstoModel(rows)
}

func (r *taskRepository) Create(ctx context.Context, task *model.Task) error {
	sqlStmt := `
	INSERT INTO tasks (id, title, priority, sort_key, created_at, updated_at, deleted_at) values(?, ?, ?, ?, ?, ?, "")
	`
	now := r.sqlite3.getNow()
	if _, err := r.sqlite3.db.Exec(sqlStmt, task.ID, task.Title, task.Priority, now, now, now); err != nil {
		return err
	}
	return nil
}

func (r *taskRepository) GetFirstByPriorityOrderBySortKeyAsc(ctx context.Context, priority float64) (*model.Task, error) {
	sqlStmt := `
	SELECT id, title, sort_key, created_at, updated_at FROM tasks WHERE priority = ? AND deleted_at == "" ORDER BY sort_key ASC LIMIT 1
	`
	row := r.sqlite3.db.QueryRow(sqlStmt, priority)
	var id, title, sort_key, create_time, update_time string
	if err := row.Scan(&id, &title, &sort_key, &create_time, &update_time); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("task not found. priority = %f: %w", priority, repositories.ErrTaskNotFound)
		}
	}
	sortKey, err := r.sqlite3.parseTime(sort_key)
	if err != nil {
		return nil, fmt.Errorf("failed to parse sort_key: %w", err)
	}
	createTime, err := r.sqlite3.parseTime(create_time)
	if err != nil {
		return nil, fmt.Errorf("failed to parse create_time: %w", err)
	}
	updateTime, err := r.sqlite3.parseTime(update_time)
	if err != nil {
		return nil, fmt.Errorf("failed to parse update_time: %w", err)
	}
	return &model.Task{
		ID:         id,
		Title:      title,
		Priority:   priority,
		SortKey:    sortKey,
		CreateTime: createTime,
		UpdateTime: updateTime,
	}, nil
}

func (r *taskRepository) Update(ctx context.Context, task *model.Task) error {
	sqlStmt := `
	UPDATE tasks SET title = ?, priority = ?, sort_key = ?, updated_at = ?, deleted_at = ? WHERE id = ?
	`
	now := r.sqlite3.getNow()
	deletedTime := ""
	if task.DeletedTime != nil {
		deletedTime = r.sqlite3.getSqlite3Time(*task.DeletedTime).String()
	}
	if _, err := r.sqlite3.db.Exec(sqlStmt, task.Title, task.Priority, task.SortKey, now, deletedTime, task.ID); err != nil {
		return err
	}
	return nil
}
