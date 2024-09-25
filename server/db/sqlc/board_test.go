package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"server/utils"
	"testing"
	"time"
)

func CreateRandomBoard(t *testing.T) Board {
	name := utils.RandomString(10)
	board, err := testQueries.CreateBoard(context.Background(), name)

	assert.NoError(t, err)
	assert.NotEmpty(t, board)
	assert.Equal(t, name, board.Name)

	_, createdAtValid := utils.ConvertNullTime(board.CreatedAt)
	_, updatedAtValid := utils.ConvertNullTime(board.UpdatedAt)

	assert.True(t, createdAtValid)
	assert.True(t, updatedAtValid)
	assert.NotZero(t, board.ID)

	return board
}

func TestCreateBoard(t *testing.T) {
	CreateRandomBoard(t)
}

func TestGetBoard(t *testing.T) {
	board := CreateRandomBoard(t)
	actualBoard, err := testQueries.GetBoard(context.Background(), board.ID)

	assert.NoError(t, err)
	assert.NotEmpty(t, actualBoard)
	assert.Equal(t, board.ID, actualBoard.ID)
	assert.Equal(t, board.Name, actualBoard.Name)

	createdAt, createdAtValid := utils.ConvertNullTime(board.CreatedAt)
	actualCreatedAt, actualCreatedAtValid := utils.ConvertNullTime(actualBoard.CreatedAt)

	updatedAt, updatedAtValid := utils.ConvertNullTime(board.UpdatedAt)
	actualUpdatedAt, actualUpdatedAtValid := utils.ConvertNullTime(actualBoard.UpdatedAt)

	if createdAtValid && actualCreatedAtValid {
		assert.WithinDuration(t, createdAt, actualCreatedAt, time.Second)
	}

	if updatedAtValid && actualUpdatedAtValid {
		assert.WithinDuration(t, updatedAt, actualUpdatedAt, time.Second)
	}
}

func TestDeleteBoard(t *testing.T) {
	board := CreateRandomBoard(t)
	deletedBoard, err := testQueries.DeleteBoard(context.Background(), board.ID)

	assert.NoError(t, err)
	assert.Equal(t, board.ID, deletedBoard.ID)
	assert.Equal(t, board.Name, deletedBoard.Name)

	createdAt, createdAtValid := utils.ConvertNullTime(board.CreatedAt)
	deletedCreatedAt, deletedCreatedAtValid := utils.ConvertNullTime(deletedBoard.CreatedAt)

	updatedAt, updatedAtValid := utils.ConvertNullTime(board.UpdatedAt)
	deletedUpdatedAt, deletedUpdatedAtValid := utils.ConvertNullTime(deletedBoard.UpdatedAt)

	if createdAtValid && deletedCreatedAtValid {
		assert.WithinDuration(t, createdAt, deletedCreatedAt, time.Second)
	}

	if updatedAtValid && deletedUpdatedAtValid {
		assert.WithinDuration(t, updatedAt, deletedUpdatedAt, time.Second)
	}

	actualBoard, err := testQueries.GetBoard(context.Background(), board.ID)
	assert.Error(t, err)
	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Empty(t, actualBoard)
}

func TestUpdateBoard(t *testing.T) {
	board := CreateRandomBoard(t)

	args := UpdateBoardParams{
		Name: utils.RandomString(16),
		ID:   board.ID,
	}

	updatedBoard, err := testQueries.UpdateBoard(context.Background(), args)

	assert.NoError(t, err)
	assert.NotEmpty(t, updatedBoard)
	assert.Equal(t, board.ID, updatedBoard.ID)
	assert.Equal(t, args.Name, updatedBoard.Name)

	createdAt, createdAtValid := utils.ConvertNullTime(board.CreatedAt)
	updatedCreatedAt, updatedCreatedAtValid := utils.ConvertNullTime(updatedBoard.CreatedAt)

	updatedAt, updatedAtValid := utils.ConvertNullTime(updatedBoard.UpdatedAt)

	if createdAtValid && updatedCreatedAtValid {
		assert.WithinDuration(t, createdAt, updatedCreatedAt, time.Second)
	}

	if updatedAtValid {
		assert.WithinDuration(t, time.Now(), updatedAt, time.Second)
	}
}

func TestListBoards(t *testing.T) {
	// Create multiple boards for testing
	for i := 0; i < 10; i++ {
		CreateRandomBoard(t)
	}

	args := ListBoardsParams{
		Limit:  5,
		Offset: 0,
	}

	boards, err := testQueries.ListBoards(context.Background(), args)

	assert.NoError(t, err)
	assert.Len(t, boards, 5)

	for _, board := range boards {
		assert.NotEmpty(t, board.ID)
		assert.NotEmpty(t, board.Name)

		_, createdAtValid := utils.ConvertNullTime(board.CreatedAt)
		_, updatedAtValid := utils.ConvertNullTime(board.UpdatedAt)

		assert.True(t, createdAtValid)
		assert.True(t, updatedAtValid)
	}
}
