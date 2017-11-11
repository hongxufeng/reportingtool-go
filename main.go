package main

import (
	"fmt"
	"utils/service"
	"utils/config"
	"os"
	"module/user"
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
	//fmt.Printf("\n%v\n\n", conf)
	server, err := service.New(&conf,false,false,false)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	server.AddModule("user",&user.UserModule{})
	fmt.Println(server.StartService())
}

