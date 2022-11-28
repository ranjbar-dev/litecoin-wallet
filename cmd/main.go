package main

import (
	"fmt"
	"github.com/ranjbar-dev/litecoin-wallet/enums"
)

func main() {

	node := enums.TEST_NODE

	w, _ := litecoinWallet.CreateLitecoinWallet(node, "ea7237b66dc3b913eb80ef94a1e9dfe6ee6843413299c380fb58c90a665b1813")

	a := &litecoinWallet.Crawler{
		Node: enums.TEST_NODE,
		Addresses: []string{
			w.Address,
		},
	}

	fmt.Println(a.ScanBlocks(30))
}
