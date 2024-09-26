package db

import (
	"context"
	"database/sql"
	"math"
	"server/utils"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomBoard(t *testing.T) (int32, error) {
	board, err := testQueries.CreateBoard(context.Background(), utils.RandomString(10))
	require.NoError(t, err)

	columns := []CreateColumnParams{
		{Name: utils.RandomString(10), Position: 0, BoardID: sql.NullInt32{Int32: board.ID, Valid: true}},
		{Name: utils.RandomString(10), Position: 1, BoardID: sql.NullInt32{Int32: board.ID, Valid: true}},
	}

	for _, col := range columns {
		_, err := testQueries.CreateColumn(context.Background(), col)
		require.NoError(t, err)
	}

	tasks := []CreateTaskParams{
		{Title: utils.RandomString(10), Description: sql.NullString{String: utils.RandomString(10), Valid: true}, ColumnID: sql.NullInt32{Int32: 1, Valid: true}, Position: 0},
		{Title: utils.RandomString(10), Description: sql.NullString{String: utils.RandomString(10), Valid: true}, ColumnID: sql.NullInt32{Int32: 2, Valid: true}, Position: 1},
	}

	for _, task := range tasks {
		_, err := testQueries.CreateTask(context.Background(), task)
		require.NoError(t, err)
	}

	return board.ID, nil
}

func TestUpdateBoardTx(t *testing.T) {
	store := NewStore(testDbContext)
	boardID, err := createRandomBoard(t)
	require.NoError(t, err)

	// Retrieve existing columns for the board
	columnsOrigin, err := testQueries.ListColumns(context.Background(), ListColumnsParams{
		BoardID: sql.NullInt32{Int32: boardID, Valid: true},
		Offset:  0,
		Limit:   math.MaxInt32,
	})
	require.NoError(t, err)
	require.NoError(t, err)

	// Prepare the updated column and task (card) data
	columns := []ColumnRequest{
		{
			ID:    columnsOrigin[0].ID,
			Title: "Updated Column 1",
			Cards: []CardRequest{
				{ID: 1, Title: "Updated Task 1", Description: "Updated Description 1"},
			},
		},
		{
			ID:    columnsOrigin[1].ID, // Use existing column ID
			Title: "Updated Column 2",
			Cards: []CardRequest{
				{ID: 0, Title: "New Task 2", Description: "New Description 2"}, // New task
			},
		},
	}

	args := UpdateBoardTxParams{
		Column:  columns,
		BoardID: boardID,
	}

	result, err := store.UpdateBoardTx(context.Background(), args)
	require.NoError(t, err)
	assert.Equal(t, "update success", result)

	updatedColumns, err := testQueries.ListColumns(context.Background(), ListColumnsParams{
		BoardID: sql.NullInt32{Int32: boardID, Valid: true},
		Offset:  0,
		Limit:   math.MaxInt32,
	})
	require.NoError(t, err)
	assert.Equal(t, 2, len(updatedColumns))
	assert.Equal(t, "Updated Column 1", updatedColumns[0].Name)
	assert.Equal(t, "Updated Column 2", updatedColumns[1].Name)

	updatedTasksCol1, err := testQueries.ListTasks(context.Background(), ListTasksParams{
		ColumnID: sql.NullInt32{Int32: updatedColumns[0].ID, Valid: true},
		Offset:   0,
		Limit:    math.MaxInt32,
	})
	require.NoError(t, err)
	assert.Equal(t, 1, len(updatedTasksCol1))
	assert.Equal(t, "Updated Task 1", updatedTasksCol1[0].Title)
	assert.Equal(t, "Updated Description 1", updatedTasksCol1[0].Description.String)

	updatedTasksCol2, err := testQueries.ListTasks(context.Background(), ListTasksParams{
		ColumnID: sql.NullInt32{Int32: updatedColumns[1].ID, Valid: true},
		Offset:   0,
		Limit:    math.MaxInt32,
	})
	require.NoError(t, err)
	assert.Equal(t, 1, len(updatedTasksCol2))
	assert.Equal(t, "New Task 2", updatedTasksCol2[0].Title)
	assert.Equal(t, "New Description 2", updatedTasksCol2[0].Description.String)
}

// TestUpdateBoardTx_Error tests the UpdateBoardTx function with invalid inputs to verify error handling
func TestUpdateBoardTx_Error(t *testing.T) {
	store := NewStore(testDbContext)
	boardID, err := createRandomBoard(t)
	require.NoError(t, err)

	args := UpdateBoardTxParams{
		Column:  []ColumnRequest{{ID: 999, Title: "Nonexistent Column"}},
		BoardID: boardID,
	}

	result, err := store.UpdateBoardTx(context.Background(), args)
	assert.Error(t, err)
	assert.Empty(t, result)
}
