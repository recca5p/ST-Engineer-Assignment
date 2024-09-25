package db

import (
	"context"
	"database/sql"
	"server/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTask(t *testing.T) Task {
	arg := CreateTaskParams{
		Title:       utils.RandomString(10),
		Description: sql.NullString{String: utils.RandomString(32), Valid: true},
		ColumnID:    sql.NullInt32{Int32: 1, Valid: true},
		Position:    1,
		DueDate:     sql.NullTime{Time: time.Now().AddDate(0, 1, 0), Valid: true},
	}

	task, err := testQueries.CreateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task)

	require.Equal(t, arg.Title, task.Title)
	require.Equal(t, arg.Description, task.Description)
	require.Equal(t, arg.ColumnID, task.ColumnID)
	require.Equal(t, arg.Position, task.Position)
	require.True(t, task.DueDate.Valid)

	require.NotZero(t, task.ID)
	require.NotZero(t, task.CreatedAt)
	require.NotZero(t, task.UpdatedAt)

	return task
}

func TestCreateTask(t *testing.T) {
	createRandomTask(t)
}

func TestDeleteTask(t *testing.T) {
	task := createRandomTask(t)

	deletedTask, err := testQueries.DeleteTask(context.Background(), task.ID)
	require.NoError(t, err)
	require.NotEmpty(t, deletedTask)

	require.Equal(t, task.ID, deletedTask.ID)
	require.Equal(t, task.Title, deletedTask.Title)
	require.Equal(t, task.Description, deletedTask.Description)
	require.Equal(t, task.ColumnID, deletedTask.ColumnID)
	require.Equal(t, task.Position, deletedTask.Position)
	require.Equal(t, task.DueDate, deletedTask.DueDate)
}

func TestGetTask(t *testing.T) {
	task := createRandomTask(t)

	fetchedTask, err := testQueries.GetTask(context.Background(), task.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedTask)

	require.Equal(t, task.ID, fetchedTask.ID)
	require.Equal(t, task.Title, fetchedTask.Title)
	require.Equal(t, task.Description, fetchedTask.Description)
	require.Equal(t, task.ColumnID, fetchedTask.ColumnID)
	require.Equal(t, task.Position, fetchedTask.Position)
	require.Equal(t, task.DueDate, fetchedTask.DueDate)
}

func TestListTasks(t *testing.T) {
	var lastTask Task
	for i := 0; i < 10; i++ {
		lastTask = createRandomTask(t)
	}

	arg := ListTasksParams{
		ColumnID: lastTask.ColumnID,
		Limit:    5,
		Offset:   0,
	}

	tasks, err := testQueries.ListTasks(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, tasks)

	for _, task := range tasks {
		require.NotEmpty(t, task)
		require.Equal(t, lastTask.ColumnID, task.ColumnID)
	}
}

func TestUpdateTask(t *testing.T) {
	task := createRandomTask(t)

	arg := UpdateTaskParams{
		Title:       utils.RandomString(16),
		Description: sql.NullString{String: utils.RandomString(32), Valid: true},
		Position:    task.Position + 1,
		DueDate:     sql.NullTime{Time: time.Now().AddDate(0, 2, 0), Valid: true}, // due date two months from now
		ID:          task.ID,
	}

	updatedTask, err := testQueries.UpdateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedTask)

	require.Equal(t, task.ID, updatedTask.ID)
	require.Equal(t, arg.Title, updatedTask.Title)
	require.Equal(t, arg.Description, updatedTask.Description)
	require.Equal(t, arg.Position, updatedTask.Position)
}
