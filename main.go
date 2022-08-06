package main

import (
	"github.com/labstack/echo"
)

func main() {
	// CONNECT TO MYSQL DATABASE
	connect()

	e := echo.New()
	e.GET("/get-all", getAll)
	e.GET("/get-by-id/:id", getById)
	e.GET("/get-by-task/:task", getByTask)
	e.POST("/create", create)
	e.PUT("/update", update)
	e.DELETE("/delete/:id", delete)
	e.Logger.Fatal(e.Start(":1323"))
}
