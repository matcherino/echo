package echo

import "net/http"

// WrapMiddleware exposes the internal wrapMiddleware function.
func WrapMiddleware(mw interface{}) MiddlewareFunc {
	if _, ok := mw.(func(*Context, http.ResponseWriter, *http.Request)); ok {
		return wrapMiddleware(WrapHandler(mw))
	}

	return wrapMiddleware(mw)
}

// WrapHandler exposes the internal  wrapHandler function.
func WrapHandler(h interface{}) HandlerFunc {
	// fmt.Printf("echo.WrapHandler type of handler %v\n", reflect.TypeOf(h))

	// if _, ok := h.([]interface{}); ok {
	// 	panic("echo.WrapHandler wrong interface type")
	// }

	if ch, ok := h.(func(*Context, http.ResponseWriter, *http.Request)); ok {
		return func(c *Context) error {
			ch(c, c.response, c.request)
			return nil
		}
	}

	return wrapHandler(h)
}
