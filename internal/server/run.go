package server

import "net/http"

func Run(addr string) error {
	sv := NewServer()

	http.HandleFunc("/todos", sv.GetTodoList)

	return http.ListenAndServe(addr, nil)
}
