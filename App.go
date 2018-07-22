package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
	. "./backend"
)

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("frontend", true)))
	backend.V1(r.Group("/api/v1"))
	r.Run(":3000")
}

