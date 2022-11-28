package enums

import "github.com/ranjbar-dev/litecoin-wallet/blockDaemon"

type Node struct {
	Config blockDaemon.ConfigBlockDaemon
	Test   bool
}

var (
	MAIN_NODE = Node{
		Config: blockDaemon.MainNet,
		Test:   false,
	}
	TEST_NODE = Node{
		Config: blockDaemon.TestNet,
		Test:   true,
	}
)
