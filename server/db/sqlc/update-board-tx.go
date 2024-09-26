package db

import (
	"context"
	"database/sql"
	"math"
)

type CardRequest struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ColumnRequest struct {
	ID    int32         `json:"id"`
	Title string        `json:"title"`
	Cards []CardRequest `json:"cards"`
}

type UpdateBoardTxParams struct {
	Column  []ColumnRequest `json:"boards"`
	BoardID int32           `json:"id"`
}

func (store *SQLStore) UpdateBoardTx(ctx context.Context, args UpdateBoardTxParams) (string, error) {
	var result string

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		for columnIndex, column := range args.Column {
			existingColumn, err := q.GetColumn(ctx, column.ID)
			if err != nil {
				return err
			}

			if &existingColumn != nil {
				updateExistingColumn := UpdateColumnParams{
					Name:     column.Title,
					ID:       column.ID,
					Position: int32(columnIndex),
				}
				_, err = q.UpdateColumn(ctx, updateExistingColumn)
				if err != nil {
					return err
				}
			} else {
				newColumn := CreateColumnParams{
					Name:     column.Title,
					Position: int32(columnIndex),
					BoardID:  sql.NullInt32{Int32: args.BoardID, Valid: true},
				}
				_, err = q.CreateColumn(ctx, newColumn)
				if err != nil {
					return err
				}
			}

			listTasksParams := ListTasksParams{
				ColumnID: sql.NullInt32{Int32: column.ID, Valid: true},
				Offset:   0,
				Limit:    math.MaxInt32,
			}
			existingTasks, err := q.ListTasks(ctx, listTasksParams)
			if err != nil {
				return err
			}

			taskMap := make(map[int32]Task)
			for _, task := range existingTasks {
				taskMap[task.ID] = task
			}

			for taskIndex, task := range column.Cards {
				if existingTask, ok := taskMap[task.ID]; ok {
					existingTask.Title = task.Title
					existingTask.Description = sql.NullString{String: task.Description, Valid: true}
					existingTask.Position = int32(taskIndex)
					updateTaskParams := UpdateTaskParams{
						Title:       existingTask.Title,
						ID:          existingTask.ID,
						Position:    existingTask.Position,
						Description: sql.NullString{String: task.Description, Valid: true},
					}
					_, err = q.UpdateTask(ctx, updateTaskParams)
					if err != nil {
						return err
					}
					delete(taskMap, updateTaskParams.ID)
				} else {
					newTask := CreateTaskParams{
						Title:       task.Title,
						Description: sql.NullString{String: task.Description, Valid: true},
						ColumnID:    sql.NullInt32{Int32: column.ID, Valid: true},
						Position:    int32(taskIndex),
					}
					_, err = q.CreateTask(ctx, newTask)
					if err != nil {
						return err
					}
				}
			}

			for taskID := range taskMap {
				_, err = q.DeleteTask(ctx, taskID)
				if err != nil {
					return err
				}
			}
		}
		return err
	})

	if err != nil {
		return "", err
	}
	result = "update success"

	return result, err
}
