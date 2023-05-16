package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateTodo(t *testing.T) {
	createRamdomTodo(t, "CreateTodo Test")
}

func TestGetTodoByID(t *testing.T) {
	todo1 := createRamdomTodo(t, "GetTodoByID Test")
	todo2, err := testQueries.GetTodo(context.Background(), todo1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, todo2)

	require.Equal(t, todo1.ID, todo2.ID)
	require.Equal(t, todo1.Title, todo2.Title)
	require.Equal(t, todo1.Note, todo2.Note)
	require.WithinDuration(t, todo1.DueDate, todo2.DueDate, time.Second)

}

func TestListTodos(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRamdomTodo(t, fmt.Sprintf("ListTodo Test %d", i))
	}

	arg := ListTodosParams{
		Limit:  5,
		Offset: 0,
	}

	todos, err := testQueries.ListTodos(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, todos)

	for _, todo := range todos {
		require.NotEmpty(t, todo)
	}

}

func TestUpdateTodo(t *testing.T) {
	todo1 := createRamdomTodo(t, "UpdateTodo Test")

	arg := UpdateTodoParams{
		ID:      todo1.ID,
		Title:   "Updated " + todo1.Title,
		Note:    "Updated " + todo1.Note,
		DueDate: time.Now(),
	}

	todo2, err := testQueries.UpdateTodo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, todo2)

	require.Equal(t, todo1.ID, todo2.ID)
	require.Equal(t, arg.Title, todo2.Title)
	require.Equal(t, arg.Note, todo2.Note)
	require.WithinDuration(t, arg.DueDate, todo2.DueDate, time.Second)
}

func TestDeleteTodo(t *testing.T) {
	todo1 := createRamdomTodo(t, "DeleteTodo")
	err := testQueries.DeleteTodo(context.Background(), todo1.ID)
	require.NoError(t, err)

	todo2, err := testQueries.GetTodo(context.Background(), todo1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, todo2)
}

func createRamdomTodo(t *testing.T, name string) Todo {
	currentTime := time.Now()

	arg := CreateTodoParams{
		Title:   "Createed By " + name,
		Note:    "Do something " + name,
		DueDate: currentTime,
	}

	todo, err := testQueries.CreateTodo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, todo)

	require.Equal(t, arg.Title, todo.Title)
	require.Equal(t, arg.Note, todo.Note)

	require.NotZero(t, todo.ID)
	require.NotZero(t, todo.CreatedAt)
	require.NotZero(t, todo.DueDate)
	require.WithinDuration(t, arg.DueDate, todo.DueDate, time.Second)

	return todo
}
