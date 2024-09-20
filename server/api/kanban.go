package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	db "server/db/sqlc"
	"server/utils"
)

type KanbanBoard struct {
	ID        int32          `json:"id"`
	Name      string         `json:"name"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	Columns   []KanbanColumn `json:"columns"`
}

type KanbanColumn struct {
	ID        int32        `json:"id"`
	Name      string       `json:"name"`
	Position  int64        `json:"position"`
	CreatedAt string       `json:"created_at"`
	UpdatedAt string       `json:"updated_at"`
	Tasks     []KanbanTask `json:"tasks"`
}

type KanbanTask struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Position    int32  `json:"position"`
	DueDate     string `json:"due_date"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type listKanbanRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5"`
}

func (server *Server) GetKanbanBoards(ctx *gin.Context) {
	var req listKanbanRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.ListBoardsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	boards, err := server.store.ListBoards(ctx, args)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, boards)
}

type kanbanRequest struct {
	ID int32 `uri:"id"`
}

func (server *Server) GetKanbanBoard(ctx *gin.Context) {
	var req kanbanRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	board, err := server.store.GetBoard(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	kanbanBoard := KanbanBoard{
		ID:        board.ID,
		Name:      board.Name,
		CreatedAt: utils.FormatNullTime(board.CreatedAt),
		UpdatedAt: utils.FormatNullTime(board.UpdatedAt),
	}

	colArgs := db.ListColumnsParams{
		Limit:   math.MaxInt32,
		Offset:  0,
		BoardID: sql.NullInt32{Valid: true, Int32: board.ID},
	}

	columns, err := server.store.ListColumns(ctx, colArgs)

	for _, column := range columns {
		taskArgs := db.ListTasksParams{
			Limit:    math.MaxInt32,
			Offset:   0,
			ColumnID: sql.NullInt32{Valid: true, Int32: column.ID},
		}
		tasks, err := server.store.ListTasks(ctx, taskArgs)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		// Initialize KanbanColumn and add tasks
		kanbanColumn := KanbanColumn{
			ID:        column.ID,
			Name:      column.Name,
			Position:  int64(column.Position),
			CreatedAt: utils.FormatNullTime(column.CreatedAt),
			UpdatedAt: utils.FormatNullTime(column.UpdatedAt),
		}

		for _, task := range tasks {
			kanbanTask := KanbanTask{
				ID:          task.ID,
				Title:       task.Title,
				Description: task.Description.String,
				Position:    task.Position,
				DueDate:     utils.FormatNullTime(task.DueDate),
				CreatedAt:   utils.FormatNullTime(column.CreatedAt),
				UpdatedAt:   utils.FormatNullTime(column.UpdatedAt),
			}
			kanbanColumn.Tasks = append(kanbanColumn.Tasks, kanbanTask)
		}

		kanbanBoard.Columns = append(kanbanBoard.Columns, kanbanColumn)
	}

	ctx.JSON(http.StatusOK, kanbanBoard)
}
