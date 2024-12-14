package api

import (
	"errors"
	"net/http"

	"github.com/ValeryBMSTU/web-10/pkg/vars"

	"github.com/labstack/echo/v4"
)

// GetHello возвращает случайное приветствие пользователю
func (srv *Server) GetQuery(e echo.Context) error {
	msg, err := srv.uc.FetchQuery()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, msg)
}

// PostHello Помещает новый вариант приветствия в БД
func (srv *Server) PostQuery(e echo.Context) error {
	input := struct {
		Msg *string `json:"msg"`
	}{}

	err := e.Bind(&input)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	if input.Msg == nil {
		return e.String(http.StatusBadRequest, "msg is empty")
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
