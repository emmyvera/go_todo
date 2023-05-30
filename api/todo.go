package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/emmyvera/go_todo/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Describe how we want our request body to be
// It also contains the request validations
type cteateTodoRequest struct {
	Title   string    `json:"title" binding:"required"`
	Note    string    `json:"note" binding:"required"`
	DueDate time.Time `json:"due_date" binding:"required"`
}

// To handle the POST request to create a new todo
func (server *Server) createTodo(ctx *gin.Context) {
	var req cteateTodoRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTodoParams{
		Title:   req.Title,
		Note:    req.Note,
		DueDate: req.DueDate,
	}

	todo, err := server.store.CreateTodo(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, todo)

}

type getTodoRequest struct {
	ID int64 `url:"id" binding:"required,min=1"`
}

// To handle the GET request to get a specifix todo
func (server *Server) getTodo(ctx *gin.Context) {
	var req getTodoRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	todo, err := server.store.GetTodo(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, todo)
}

type listTodoRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// To handle the GET request to get a all todo
// Also has pagination functionality
func (server *Server) listTodo(ctx *gin.Context) {
	var req listTodoRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListTodosParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	todos, err := server.store.ListTodos(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, todos)
}

type updateTodoRequestUri struct {
	ID int64 `url:"id" binding:"required,min=1"`
}

type updateTodoRequestBody struct {
	Title   string    `json:"title" binding:"required"`
	Note    string    `json:"note" binding:"required"`
	DueDate time.Time `json:"due_date" binding:"required"`
}

// To handle the PUT request to update a specifix todo
func (server *Server) updateTodo(ctx *gin.Context) {
	var req updateTodoRequestUri
	var reqBody updateTodoRequestBody
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateTodoParams{
		ID:      req.ID,
		Title:   reqBody.Title,
		Note:    reqBody.Note,
		DueDate: reqBody.DueDate,
	}

	todo, err := server.store.UpdateTodo(ctx, arg)
	if err != nil {

		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, todo)

}

type delTodoRequest struct {
	ID int64 `url:"id" binding:"required,min=1"`
}

// To handle the DELETE request to delete a specifix todo
func (server *Server) delTodo(ctx *gin.Context) {
	var req delTodoRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteTodo(ctx, req.ID)
	if err != nil {

		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON((http.StatusInternalServerError), errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "Todo Deleted")
}
