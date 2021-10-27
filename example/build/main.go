package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
    "os"
)

var (
    gitHash    string
    buildStamp string
    goVersion  string
)

func main() {
    args := os.Args
    if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
        fmt.Println(goVersion)
        fmt.Printf("Git Commit Hash: %s\n", gitHash)
        fmt.Printf("UTC Build Time : %s\n", buildStamp)
        return
    }

    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, "Hello, welcome to gin world!")
    })
    err := r.Run(":80")
    if err != nil {
        panic(err)
    }
}
