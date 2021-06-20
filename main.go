package main

import (
	"flag"
	"odd-number-service/handler/router"

	"github.com/gin-gonic/gin"
)

var (
	port string
)

func init() {
	flag.StringVar(&port, "port", ":6060", "this is a REST server port")
}

func main() {
	flag.Parse()
	server := gin.Default()
	router.AttachRouter(server)
	server.Run(port)

}
