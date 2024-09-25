package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"server/utils"
	"testing"
	"time"
)

func CreateRandomColumn(t *testing.T) Column {
	args := CreateColumnParams{
		Name:     utils.RandomString(10),
		BoardID:  sql.NullInt32{Int32: 1, Valid: true},
		Position: 1,
	}

	column, err := testQueries.CreateColumn(context.Background(), args)

	assert.NoError(t, err)
	assert.NotEmpty(t, column)
	assert.Equal(t, args.Name, column.Name)
	assert.Equal(t, args.BoardID.Int32, column.BoardID.Int32)
	assert.Equal(t, args.Position, column.Position)

	_, createdAtValid := utils.ConvertNullTime(column.CreatedAt)
	_, updatedAtValid := utils.ConvertNullTime(column.UpdatedAt)

	assert.True(t, createdAtValid)
	assert.True(t, updatedAtValid)
	assert.NotZero(t, column.ID)

	return column
}

func TestCreateColumn(t *testing.T) {
	CreateRandomColumn(t)
}

func TestGetColumn(t *testing.T) {
	column := CreateRandomColumn(t)
	actualColumn, err := testQueries.GetColumn(context.Background(), column.ID)

	assert.NoError(t, err)
	assert.NotEmpty(t, actualColumn)
	assert.Equal(t, column.ID, actualColumn.ID)
	assert.Equal(t, column.Name, actualColumn.Name)

	createdAt, createdAtValid := utils.ConvertNullTime(column.CreatedAt)
	actualCreatedAt, actualCreatedAtValid := utils.ConvertNullTime(actualColumn.CreatedAt)

	updatedAt, updatedAtValid := utils.ConvertNullTime(column.UpdatedAt)
	actualUpdatedAt, actualUpdatedAtValid := utils.ConvertNullTime(actualColumn.UpdatedAt)

	if createdAtValid && actualCreatedAtValid {
		assert.WithinDuration(t, createdAt, actualCreatedAt, time.Second)
	}

	if updatedAtValid && actualUpdatedAtValid {
		assert.WithinDuration(t, updatedAt, actualUpdatedAt, time.Second)
	}
}

func TestDeleteColumn(t *testing.T) {
	column := CreateRandomColumn(t)
	deletedColumn, err := testQueries.DeleteColumn(context.Background(), column.ID)

	assert.NoError(t, err)
	assert.Equal(t, column.ID, deletedColumn.ID)
	assert.Equal(t, column.Name, deletedColumn.Name)

	createdAt, createdAtValid := utils.ConvertNullTime(column.CreatedAt)
	deletedCreatedAt, deletedCreatedAtValid := utils.ConvertNullTime(deletedColumn.CreatedAt)

	updatedAt, updatedAtValid := utils.ConvertNullTime(column.UpdatedAt)
	deletedUpdatedAt, deletedUpdatedAtValid := utils.ConvertNullTime(deletedColumn.UpdatedAt)

	if createdAtValid && deletedCreatedAtValid {
		assert.WithinDuration(t, createdAt, deletedCreatedAt, time.Second)
	}

	if updatedAtValid && deletedUpdatedAtValid {
		assert.WithinDuration(t, updatedAt, deletedUpdatedAt, time.Second)
	}

	// Check if the column is indeed deleted
	actualColumn, err := testQueries.GetColumn(context.Background(), column.ID)
	assert.Error(t, err)
	assert.EqualError(t, err, sql.ErrNoRows.Error())
	assert.Empty(t, actualColumn)
}

func TestUpdateColumn(t *testing.T) {
	column := CreateRandomColumn(t)

	args := UpdateColumnParams{
		Name:     "Updated Test Column",
		Position: 2,
		ID:       column.ID,
	}

	updatedColumn, err := testQueries.UpdateColumn(context.Background(), args)

	assert.NoError(t, err)
	assert.NotEmpty(t, updatedColumn)
	assert.Equal(t, column.ID, updatedColumn.ID)
	assert.Equal(t, args.Name, updatedColumn.Name)
	assert.Equal(t, args.Position, updatedColumn.Position)

	createdAt, createdAtValid := utils.ConvertNullTime(column.CreatedAt)
	updatedCreatedAt, updatedCreatedAtValid := utils.ConvertNullTime(updatedColumn.CreatedAt)

	updatedAt, updatedAtValid := utils.ConvertNullTime(updatedColumn.UpdatedAt)

	if createdAtValid && updatedCreatedAtValid {
		assert.WithinDuration(t, createdAt, updatedCreatedAt, time.Second)
	}

	if updatedAtValid {
		assert.WithinDuration(t, time.Now(), updatedAt, time.Second)
	}
}
func TestListColumns(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomColumn(t)
	}

	args := ListColumnsParams{
		BoardID: sql.NullInt32{Int32: 1, Valid: true}, // Assuming 1 is a valid board ID
		Limit:   5,
		Offset:  0,
	}

	columns, err := testQueries.ListColumns(context.Background(), args)

	assert.NoError(t, err)
	assert.Len(t, columns, 5)

	for _, column := range columns {
		assert.NotEmpty(t, column.ID)
		assert.NotEmpty(t, column.Name)

		_, createdAtValid := utils.ConvertNullTime(column.CreatedAt)
		_, updatedAtValid := utils.ConvertNullTime(column.UpdatedAt)

		assert.True(t, createdAtValid)
		assert.True(t, updatedAtValid)
	}
}
