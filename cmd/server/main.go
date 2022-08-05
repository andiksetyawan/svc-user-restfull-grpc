package main

import (
	"flag"

	"svc-user/internal/app"
)

func main() {
	port := flag.String("port", "8080", "setup port server")
	flag.Parse()
	app.Run(*port)
}
