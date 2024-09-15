package main

import (
	"github.com/gin-gonic/gin"
	"zip-like/handlers"
)

func main() {
	r := gin.Default()
	r.POST("/compress", handlers.CompressHandler)
	r.POST("/decompress", handlers.DecompressHandler)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
