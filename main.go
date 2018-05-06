package main

import (
    "github.com/kawabatas/go-crud-sample/controller"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // API namespace
    v1 := router.Group("/api/v1")
    {
        v1.GET("/tasks", controller.TasksGET)
    }

    router.GET("/", controller.IndexGET)
    router.Run(":8080")
}
