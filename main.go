package main

import (
	"fmt"
	"gin-demo/pkg/setting"
	"gin-demo/routers"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"net/http"
)

func main() {

	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf("127.0.0.1:%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}