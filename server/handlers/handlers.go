package handlers

import (
	"echo-restapi/pkg"
	"echo-restapi/service"
	"echo-restapi/types"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

type Handlers struct {
	srv *service.Service
}

func NewHandlers(users *service.Service) *Handlers {
	return &Handlers{
		srv: users,
	}
}

func (h *Handlers) SaveUser(c echo.Context) error {
	user := new(types.User)
	id, oldusers, err := h.srv.CheckJsonFile("json.json")
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Файл еще не создан SaveUser"))
		return err
	}
	user.Id = id
	if err = c.Bind(&user); err != nil {
		pkg.LogError(errors.Wrap(err, "Ошибка кодирования SaveUser"))
		return err
	}
	if user.Name == "" {
		return c.HTML(http.StatusOK, "Некорректные входные данные")
	}
	users, err := h.srv.SaveUser(user, oldusers)
	if err != nil {
		return c.HTML(http.StatusOK, "Не удалось сохранить пользователя")
	}
	return c.JSONPretty(http.StatusCreated, users, "")
}

func (h *Handlers) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Err with strconv.Atoi(id) in GetUser"))
		return err
	}
	user, err := h.srv.GetUser(id)
	if err != nil {
		return c.HTML(http.StatusOK, "Не удалось получить инфо о пользователе")
	}
	if user.Id == 0 {
		return c.HTML(http.StatusOK, "Неверный id пользователя")
	}
	return c.JSONPretty(http.StatusCreated, user, "")
}

func (h *Handlers) GetUsers(c echo.Context) error {
	users, err := h.srv.GetUsers()
	if err != nil {
		return c.HTML(http.StatusOK, "Не удалось получить инфо о пользователях")
	}
	return c.JSONPretty(http.StatusCreated, users, "")
}

func (h *Handlers) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Err with strconv.Atoi(id) in UpdateUser"))
		return err
	}
	req := c.Request()
	name := req.FormValue("name")
	if name == "" {
		return c.HTML(http.StatusOK, "Некорректное имя пользователя")
	}
	okId, err := h.srv.UpdateUser(id, name)
	if err != nil {
		return c.HTML(http.StatusBadGateway, "Не удалось обновить пользователя")
	}
	if okId == true {
		return c.HTML(http.StatusOK, "Успешно изменено")
	} else {
		return c.HTML(http.StatusOK, "Некорректный Id пользователя")
	}
}

func (h *Handlers) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Err with strconv.Atoi(id) in DeleteUser"))
		return err
	}
	okId, err := h.srv.DeleteUser(id)
	if err != nil {
		pkg.LogError(errors.Wrap(err, "Ошибка сервиса DeleteUser"))
		return err
	}
	if okId == true {
		return c.HTML(http.StatusOK, "Успешно изменено")
	} else {
		return c.HTML(http.StatusOK, "Некорректный Id пользователя")
	}
}
