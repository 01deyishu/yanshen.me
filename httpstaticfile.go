package main

import (
	"flag"
	"fmt"
	"github.com/vishvananda/netlink"
	"net/http"
	"regexp"
	"strings"
)

func httpInfo(port string) {
	fmt.Println("输入下列链接访问静态文件服务器：")
	links, _ := netlink.LinkList()
	for _, link := range links {
		ip, _ := netlink.AddrList(link, 0)
		if ip != nil {
			url := strings.Split(ip[0].String(), "/")
			if isOk, _ := regexp.MatchString("127.*", url[0]); !isOk {
				if port == ":80" {
					fmt.Println("http://" + url[0])
				} else {
					fmt.Println("http://" + url[0] + port)
				}
			}
		}
	}
}

func main() {
	port := flag.String("port", "80", "指定监听的端口号")
	flag.Parse()
	listen := ":" + *port
	httpInfo(listen)
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(listen, nil)
	if err != nil {
		fmt.Errorf("%v", err)
	}
}
