package api

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	db "server/db/sqlc"
)

type updateColumn struct {
	ID       int32  `json:"id" binding:"required"`
	Position int32  `json:"position"`
	Name     string `json:"name"`
}

func (server *Server) UpdateColumn(ctx *gin.Context) {
	var req updateColumn
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.Position <= 0 && req.Name == "" {
		err := errors.New("Position and name is empty")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.UpdateColumnParams{
		ID:       req.ID,
		Position: req.Position,
		Name:     req.Name,
	}

	_, err := server.store.UpdateColumn(ctx, args)
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
