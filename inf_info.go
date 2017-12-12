//1.3.6.1.2.1.4.34.1.3.1.4
//1.3.6.1.2.1.2.2.1.2
package main

import (
	"fmt"
	"os/exec"
	//	"reflect"
	"strings"
)

func main() {
	response, err := exec.Command("get_ip.sh").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	ip := string(response)
	//	fmt.Println(reflect.TypeOf(ip))
	for _, j := range strings.Split(ip, "\n") {
		fmt.Println(j)
	}

	ifname, err := exec.Command("get_ifname.sh", "58").Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(ifname))
}
