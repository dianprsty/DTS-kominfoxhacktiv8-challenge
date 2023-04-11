package main

import (
	"Chap2-Challenge3/routers"
	"fmt"
)

func main() {
	port := ":8081"

	fmt.Println("Server Running at port", port)
	routers.StartServer().Run(port)
}
