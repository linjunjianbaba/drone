package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/sirupsen/logrus"
)

func main() {
  r := gin.Default()

  r.GET("/health", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H {
      version: v3
    })
  })

  if err := r.Run(":8080"); err != nil {
    logrus.WithError(err).Fatal("Couldn't listen")
  }

}