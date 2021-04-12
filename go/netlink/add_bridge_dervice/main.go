package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
)

func main() {
	la := netlink.NewLinkAttrs()
	la.Name = "foo2"
	mybridge := &netlink.Bridge{LinkAttrs: la}
	err := netlink.LinkAdd(mybridge)
	if err != nil {
		fmt.Printf("could not add %s: %v\n", la.Name, err)
	}
	eth1, err := netlink.LinkByName("enp0s31f6")
	if err != nil {
		fmt.Printf("could not link %s: %v\n", la.Name, err)
	}
	netlink.LinkSetMaster(eth1, mybridge)
}
