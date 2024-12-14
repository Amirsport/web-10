package api

import (
	"errors"
	"net/http"

	"github.com/Amirsport/web-10/pkg/vars"

	"github.com/labstack/echo/v4"
)

// GetHello возвращает случайное приветствие пользователю
func (srv *Server) GetQuery(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Name parameter is required"})
	}
	return c.String(http.StatusOK, "Hello, "+name+"!")
}

// PostHello Помещает новый вариант приветствия в БД
func (srv *Server) PostQuery(e echo.Context) error {
	input := struct {
		Msg *string `json:"name"`
	}{}

	err := e.Bind(&input)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	if input.Msg == nil {
		return e.String(http.StatusBadRequest, "name is empty")
	}

	if len([]rune(*input.Msg)) > srv.maxSize {
		return e.String(http.StatusBadRequest, "hello message too large")
	}

	err = srv.uc.SetQuery(*input.Msg)
	if err != nil {
		if errors.Is(err, vars.ErrAlreadyExist) {
			return e.String(http.StatusConflict, err.Error())
		}
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.String(http.StatusCreated, "OK")
}
