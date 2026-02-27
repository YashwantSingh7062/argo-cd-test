package controller

import (
	"fmt"
	"net/http"

	"github.com/YashwantSingh7062/argo-cd-test/lib/models"

	"github.com/labstack/echo"
)

func (c *Controller) GetUser(ctx echo.Context) error {
	user := models.User{
		Id:    "1234",
		Email: "yashwantsingh@gmail.com",
		Name:  "Yashwant Singh",
	}

	fmt.Println("user -->", user)

	return ctx.JSON(http.StatusOK, user)
}
