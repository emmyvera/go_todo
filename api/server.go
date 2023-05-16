package api

import (
	db "github.com/emmyvera/go_todo/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our todo application.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer create a new HTTP server and ssetup routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// All the list of the routes handers,
	// You can find it functionality in the todo.go
	// file locatedin the api folder
	router.POST("/todos", server.createTodo)
	router.GET("/todos/:ID", server.getTodo)
	router.GET("/todos/", server.listTodo)
	router.DELETE("/todos/:ID", server.delTodo)
	router.PUT("/todos/:ID", server.updateTodo)

	server.router = router

	return server
}

// Start runs the HTTP Server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// Handles Error Messages
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
