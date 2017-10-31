package main

import (
	"net/http"
	"strings"
	"utils"
	"fmt"
	"datahelper"
	"utils/service"
)
func mValidUser(r *http.Request, bodybytes []byte) (uid uint32, md5ok bool) {
	md5ok = true
	c, e := r.Cookie("auth")
	if e != nil {
		return 0, md5ok
	}
	auth := c.Value
	var hashcode string
	var ks []string
	if strings.Contains(auth, "%09") {
		ks = strings.Split(auth, "%09")
	} else {
		ks = strings.Split(auth, "_")
	}

	if len(ks) == 2 {
		uid, e = utils.ToUint32(ks[0])
		if e != nil {
			fmt.Println(e.Error())
		}
		hashcode = ks[1]
	}
	valid, e := datahelper.UserValid(uid, hashcode)
	if e != nil || !valid {
		return 0, md5ok
	}
	return uid, md5ok
}
func main() {
	server, err := service.New()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	server.SetmValid(mValidUser)

	fmt.Println(server.StartService())
}

