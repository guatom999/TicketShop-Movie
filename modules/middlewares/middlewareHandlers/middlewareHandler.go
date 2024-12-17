package middlewareHandlers

type (
	MiddlewareHandlerInterface interface {
	}

	middlewareHandler struct {
	}
)

func NewMiddlewareHandler() MiddlewareHandlerInterface {
	return &middlewareHandler{}
}
