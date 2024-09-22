package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	db "server/db/sqlc"
)

type createTask struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Position    int32  `json:"position" binding:"required"`
	ColumnID    int32  `json:"columnID" binding:"required"`
}

func (server *Server) CreateTask(ctx *gin.Context) {
	var req createTask
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.CreateTaskParams{
		Title:       req.Name,
		Position:    req.Position,
		Description: sql.NullString{Valid: true, String: req.Description},
		ColumnID:    sql.NullInt32{Int32: req.ColumnID, Valid: true},
	}

	_, err := server.store.CreateTask(ctx, args)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "Update success")
}

type updateTask struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Position    int32  `json:"position" binding:"required"`
	ID          int32  `json:"id" binding:"required"`
}

func (server *Server) UpdateTask(ctx *gin.Context) {
	var req updateTask
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.UpdateTaskParams{
		Title:       req.Name,
		Position:    req.Position,
		Description: sql.NullString{Valid: true, String: req.Description},
		ID:          req.ID,
	}

	_, err := server.store.UpdateTask(ctx, args)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "Update success")
}

type deleteTask struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) DeleteTask(ctx *gin.Context) {
	var req deleteTask
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err := server.store.DeleteTask(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "Update success")
}
