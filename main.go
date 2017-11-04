package main

import (
	"fmt"
	"utils/service"
)
func main() {

	server, err := service.New()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(server.StartService())
}

