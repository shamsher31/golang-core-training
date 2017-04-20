package todoapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "http://www.todoapi.com/v1"

// Client contains user credentials
type Client struct {
	Username string
	Password string
}

// Todo is todo
type Todo struct {
	ID      int    `json:"Id"`
	Content string `json:"content"`
	Done    bool   `json:"done"`
}

// NewBasicAuthClient is for authentication
func NewBasicAuthClient(username, password string) *Client {
	return &Client{
		Username: username,
		Password: password,
	}
}

// AddTodo adds todo
func (s *Client) AddTodo(todo *Todo) error {
	url := fmt.Sprint(baseURL+"/%s/todos", s.Username)
	data, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	_, err = s.doRequest(req)
	return err
}

// GetTodo returns todo item
func (s *Client) GetTodo(id int) (*Todo, error) {
	url := fmt.Sprintf(baseURL+"/%s/todo/%d", s.Username, id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}

	var data Todo
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, err
}

func (s *Client) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(s.Username, s.Password)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
