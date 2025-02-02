package main

import (
	"examplegpt.com/router"
	"github.com/gin-gonic/gin"
)

func main() {
	serv := gin.Default()

	router.Router(serv)

	serv.Run(":3030")
}
