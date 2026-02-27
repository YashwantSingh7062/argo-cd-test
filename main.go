package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/YashwantSingh7062/argo-cd-test/lib/controller"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	serviceController := controller.Controller{}
	e.GET("/", serviceController.GetUser)

	go func() {
		e.Logger.Fatal(e.Start(":8000"))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		fmt.Println("Forcefully shutting Down")
	}

	fmt.Println("ShutDown Successful")
}
