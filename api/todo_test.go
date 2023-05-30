package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/emmyvera/go_todo/db/mock"
	db "github.com/emmyvera/go_todo/db/sqlc"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetTodoAPI(t *testing.T) {
	todo := randomTodo()

	testCases := []struct {
		name          string
		todoID        int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			todoID: todo.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTodo(gomock.Any(), gomock.Eq(todo.ID)).
					Times(1).
					Return(todo, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchTodo(t, recorder.Body, todo)
			},
		},

		{
			name:   "Bad Request",
			todoID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTodo(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},

		{
			name:   "Not Found",
			todoID: todo.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTodo(gomock.Any(), gomock.Eq(todo.ID)).
					Times(1).
					Return(db.Todo{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},

		{
			name:   "Internal Found",
			todoID: todo.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetTodo(gomock.Any(), gomock.Eq(todo.ID)).
					Times(1).
					Return(db.Todo{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/todos/%d", tc.todoID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomTodo() db.Todo {
	return db.Todo{
		ID:      18,
		Title:   "Testing API Title",
		Note:    "Testing API Note",
		DueDate: time.Now(),
	}
}

func requireBodyMatchTodo(t *testing.T, body *bytes.Buffer, todo db.Todo) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotTodo db.Todo
	err = json.Unmarshal(data, &gotTodo)
	require.NoError(t, err)
	require.Equal(t, todo.Title, gotTodo.Title)
	require.Equal(t, todo.Note, gotTodo.Note)
	require.WithinDuration(t, todo.DueDate, gotTodo.DueDate, time.Second)
}
