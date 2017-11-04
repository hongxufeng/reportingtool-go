package main

import (
	"fmt"
	"utils/service"
	"utils/config"
	"os"
)

var conf config.Config

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("invalid args : %s [config]\n", os.Args[0])
		return
	}
	err := conf.LoadPath(os.Args[1])
	if err != nil {
		fmt.Println("Load Config error :", err.Error())
		return
	}
	server, err := service.New(&conf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(server.StartService())
}

