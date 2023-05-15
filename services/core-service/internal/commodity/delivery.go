package commodity

import "github.com/labstack/echo/v4"

// Handlers ...
type Handlers interface {
	List() echo.HandlerFunc
	Aggregate() echo.HandlerFunc
}
