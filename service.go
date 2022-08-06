package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func getAll(c echo.Context) error {
	db.Find(&todos)
	return c.JSON(http.StatusOK, todos)
}

func getById(c echo.Context) error {
	sql := db.Find(&todos, c.Param("id"))
	if sql.Error != nil {
		return c.JSON(http.StatusOK, sql.Error)
	}
	return c.JSON(http.StatusOK, todos)
}

func getByTask(c echo.Context) error {
	db.Find(&todos, "task = ?", c.Param("task"))
	return c.JSON(http.StatusOK, todos)
}

func create(c echo.Context) error {
	req_body := &RequestBody{}
	if err := c.Bind(&req_body); err != nil {
		return err
	}
	db.Create(&Todos{Task: req_body.Task})
	db.Last(&todos)
	return c.JSON(http.StatusOK, todos)
}

func update(c echo.Context) error {
	req_body := &RequestBody{}
	if err := c.Bind(&req_body); err != nil {
		return err
	}
	db.Model(&Todos{}).Where("id = ?", req_body.Id).Update("task", req_body.Task)
	db.Find(&todos, req_body.Id)
	return c.JSON(http.StatusOK, &todos)
}

func delete(c echo.Context) error {
	db.Find(&todos, c.Param("id"))
	if len(todos) == 0 {
		return c.String(http.StatusOK, "There is no record for id: "+c.Param("id"))

	}
	db.Delete(&Todos{}, c.Param("id"))
	return c.String(http.StatusOK, "delete success")
}
