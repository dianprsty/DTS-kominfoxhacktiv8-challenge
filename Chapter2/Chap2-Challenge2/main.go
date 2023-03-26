package main

import (
	"Chap2-Challenge2/routers"
	"fmt"
)

func main() {
	port := ":8080"

	fmt.Println("Server Running at port", port)
	routers.StartServer().Run(port)
}
