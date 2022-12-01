package test

import (
	litecoinWallet "github.com/ranjbar-dev/litecoin-wallet"
	"github.com/ranjbar-dev/litecoin-wallet/enums"
)

var node = enums.TEST_NODE
var validPrivateKey = "ea7237b66dc3b913eb80ef94a1e9dfe6ee6843413299c380fb58c90a665b1813"
var invalidPrivateKey = "invalid"
var validOwnerAddress = "tltc1qnryq8mwgtvpqx4uz5hh5wu20yvpzzd3ur3g3ae"
var invalidOwnerAddress = "tb15111190u4dz48ctn1273333ss7fmspckag341fyp0"
var validToAddress = "tltc1q39psaj2wrshycjyy7u3uytmx6sx7tyzd82zawd"
var invalidToAddress = "tb15111190u4dz48ctn1273333ss7fmspckag341fyp0"
var ltcAmount = int64(10000)

func wallet() *litecoinWallet.LitecoinWallet {
	w, _ := litecoinWallet.CreateLitecoinWallet(node, validPrivateKey)
	return w
}

func crawler() *litecoinWallet.Crawler {

	return &litecoinWallet.Crawler{
		Node:      node,
		Addresses: []string{validOwnerAddress},
	}
}
