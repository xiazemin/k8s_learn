package main

import (
	"fmt"

	"github.com/vishvananda/netlink"
)

func main() {
	la := netlink.NewLinkAttrs()
	la.Name = "foo"
	mybridge := &netlink.Bridge{LinkAttrs: la}
	err := netlink.LinkAdd(mybridge)
	if err != nil {
		fmt.Printf("could not add %s: %v\n", la.Name, err)
	}
	eth1, _ := netlink.LinkByName("eth1")
	netlink.LinkSetMaster(eth1, mybridge)
}
