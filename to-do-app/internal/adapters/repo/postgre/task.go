package repo

import (
	"context"
	"database/sql"
	"to-do-app/internal/adapters/repo/postgre/dao"
	"to-do-app/internal/model"
)

type taskRepo struct {
	table string
	db    *sql.DB
}

const tableTask = "tasks"

func NewTaskRepo(db *sql.DB) *taskRepo {
	return &taskRepo{table: tableTask, db: db}
}

func (r *taskRepo) CreateTask(ctx context.Context, task *model.Task) (*model.Task, error) {
	newTask := dao.FromDomain(task)
	query := `
		INSERT INTO ` + r.table + `
		(title, done, created_at, deadline, priority)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.QueryContext(ctx, query,
		newTask.Title,
		newTask.Done,
		newTask.CreatedAt,
		newTask.Deadline,
		newTask.Priority,
	)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *taskRepo) GetAllTask(ctx context.Context) ([]model.Task, error) {
	query := `
		SELECT * FROM ` + r.table

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	taskList := []model.Task{}
	for rows.Next() {
		task := dao.Task{}
		err = rows.Scan(
			&task.ID,
			&task.Title,
			&task.Done,
			&task.CreatedAt,
			&task.Deadline,
			&task.Priority,
		)

		if err != nil {
			return nil, err
		}

		taskList = append(taskList, dao.ToDomain(&task))
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return taskList, nil
}

func (r *taskRepo) ToggleTask(ctx context.Context, task *model.Task) error {
	object := dao.FromDomain(task)

	query := `
		UPDATE ` + r.table + `
		SET done =  $1
		WHERE id = $2
	`

	_, err := r.db.QueryContext(ctx, query,
		object.Done,
		object.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepo) DeleteTask(ctx context.Context, task *model.Task) error {
	object := dao.FromDomain(task)

	query := `DELETE FROM ` + r.table + ` WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, object.ID)
	if err != nil {
		return err
	}

	return nil
}
