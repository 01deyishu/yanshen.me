package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"net/http"
	"regexp"
	"strings"
)

func httpInfo() {
	fmt.Println("输入下列链接访问静态文件服务器：")
	links, _ := netlink.LinkList()
	for _, link := range links {
		ip, _ := netlink.AddrList(link, 0)
		if ip != nil {
			url := strings.Split(ip[0].String(), "/")
			isOk, _ := regexp.MatchString("127.*", url[0])
			if !isOk {
				fmt.Println("http://" + url[0])
			}
		}
	}
}

func main() {
	httpInfo()
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}
