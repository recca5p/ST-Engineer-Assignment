package api

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "server/db/sqlc"
	"server/token"
	"server/utils"
)

type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(cors.Default())

	/*router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/token/renew", server.renewAccessToken)*/
	router.GET("/board", server.GetKanbanBoards)
	router.GET("/board/:id", server.GetKanbanBoard)
	router.PUT("/board/:id", server.UpdateBoard)
	router.PUT("/board/update-column", server.UpdateColumn)
	router.POST("/board/create-task", server.CreateTask)
	router.PUT("/board/update-task", server.UpdateTask)
	router.DELETE("/board/delete-task/:id", server.DeleteTask)

	/*// Admin-only route
	adminRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker, []string{"admin"}))
	adminRoutes.PUT("/column", server.UpdateColumn)
	// Admin and User route
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker, []string{"admin", "user"}))
	authRoutes.GET("/board", server.GetKanbanBoards)
	authRoutes.GET("/board/:id", server.GetKanbanBoard)*/

	server.router = router
}
