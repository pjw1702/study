package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/stretchr/testify/assert"

	"testing"

	"strconv"

	"github.com/pjw1702/study/web/go/web18/model"
)

// 테스트 코드는 리팩토링 할 때 바뀐 코드의 동작이 기존 코드의 동작과 동일한지의 여부를 확인할 수 있는 중요한 요소가 된다

func TestTodos(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(MakeHandler())
	defer ts.Close()

	// /todos test
	resp, err := http.PostForm(ts.URL+"/todos", url.Values{"name": {"Test todo"}})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	var todo model.Todo
	err = json.NewDecoder(resp.Body).Decode(&todo) // 리퀘스트 Body를 언마샬링할 때 에러가 발생하는지 체크
	assert.NoError(err)
	assert.Equal(todo.Name, "Test todo")

	id1 := todo.ID
	resp, err = http.PostForm(ts.URL+"/todos", url.Values{"name": {"Test todo2"}})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal(todo.Name, "Test todo2")
	id2 := todo.ID

	// confirm ids of /todos test
	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	todos := []*model.Todo{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(len(todos), 2)
	for _, t := range todos {
		if t.ID == id1 {
			assert.Equal("Test todo", t.Name)
		} else if t.ID == id2 {
			assert.Equal("Test todo2", t.Name)
		} else {
			assert.Error(fmt.Errorf("testID should be id1 or id2"))
		}
	}

	// /complete-todo/{id} test
	resp, err = http.Get(ts.URL + "/complete-todo/" + strconv.Itoa(id1) + "?complete=true")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	// confirm contents of completed
	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	todos = []*model.Todo{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(len(todos), 2)
	for _, t := range todos {
		if t.ID == id1 {
			assert.True(t.Completed)
		}
	}

	// Delete test of /todos/{id}
	req, _ := http.NewRequest("DELETE", ts.URL+"/todos/"+strconv.Itoa(id1), nil) // http.NewRequest(): Golang에서의 http 패키지는 기본적으로 DELETE 메서드에 대한 리퀘스트 함수를 제공하지 않으므로 새로운 메서드로서의 리퀘스트 함수를 새로 생성해 주어야 한다
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	// confirm result to delete id1
	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	todos = []*model.Todo{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(len(todos), 1) // id1 is deleted
	for _, t := range todos {
		assert.Equal(t.ID, id2)
	}
}
