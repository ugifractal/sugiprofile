package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "log"
  "os"
)

var router *gin.Engine

func main() {
  port := os.Getenv("PORT")

  if port == "" {
    log.Fatal("$PORT must be set")
  }

  router = gin.Default()
  router.Use(gin.Logger())
  router.LoadHTMLGlob("templates/*")
  router.Static("/static", "static")
  router.GET("/", func(c *gin.Context){
    c.HTML(
      http.StatusOK,
      "index.html",
      gin.H{
        "title": "Home Page",
      },
    )
  })
  router.Run(":" + port)
}


