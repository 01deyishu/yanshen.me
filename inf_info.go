package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
	//	"os"
	//"github.com/alouca/gosnmp"
)

func main() {
	link := netlink.NewLinkAttrs()
	link.Name = "enp0s10"
}
