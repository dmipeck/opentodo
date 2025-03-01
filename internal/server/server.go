package server

import (
	"context"

	openapi "gitlab.build13.com/opentodo/opentodo/internal/openapi/server"
)

func NewServer() openapi.ServerInterface {
	return openapi.NewStrictHandlerWithOptions(
		&server{},
		nil,
		openapi.StrictHTTPServerOptions{
			RequestErrorHandlerFunc:  requestErrorHandler,
			ResponseErrorHandlerFunc: responseErrorHandler,
		},
	)
}

type server struct{}

func (s *server) GetTodoList(
	ctx context.Context,
	request openapi.GetTodoListRequestObject,
) (openapi.GetTodoListResponseObject, error) {
	return openapi.GetTodoList200JSONResponse([]openapi.Todo{}), nil
}
