package main

import (
    "github.com/kawabatas/go-crud-sample/controller"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/", controller.IndexGET)
    router.Run(":8080")
}
