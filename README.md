# litecoin-wallet
litecoin wallet package for creating and generating wallet, transferring LTC, getting wallet unspent transactions(UTXOs), getting wallet txs , getting wallet balance and crawling blocks to find wallet transactions
 
 ### Installation 
```
go get github.com/ranjbar-dev/litecoin-wallet@v1.0.5
```

### Supported nodes
check `enums/nodes` file  

### Wallet methods

Generating litecoin wallet
```
w := GenerateLitecoinWallet(node)
w.Address // strnig 
w.PrivateKey // strnig 
w.PublicKey // strnig 
```

Creating litecoin wallet from private key
```
w := CreateLitecoinnWallet(node,privateKeyHex)
w.Address // strnig 
w.PrivateKey // strnig 
w.PublicKey // strnig 
```

Getting wallet litecoin balance
```
balanceInSatoshi,err := w.Balance()
balanceInSatoshi // int64
```

Getting wallet UTXOs
```
utxos,err := w.UTXOs()
utxos // []response.UTXO
```

Getting wallet transactions
```
txs,err := w.Txs()
txs // []response.Transaction
```

crawl blocks for addresses transactions
```

c := &Crawler{
		Node: node, 
		Addresses: []string{
			"tb1qppv790u4dz48ctnk3p7ss7fmspckagp3wrfyp0", // list of your addresses
		},
	}
	
res, err := c.ScanBlocks(40) // scan latest 40 block on block chain and extract addressess transactions 

Example 
// *
{
    {
        "address": "tb1qppv790u4dz48ctnk3p7ss7fmspckagp3wrfyp0",
        "tranasctions": {
            {
                "tx_id": "e6160c52401949139688623ce33a6290eed43d8d564d6e16c38006c4dc28f4a8",
                "from_address": "tb1qppv790u4dz48ctnk3p7ss7fmspckagp3wrfyp0",
                "to_address": "tb1q0r23g66m9rhhak8aahsg53wfp5egt2huuc4tnu",
                "amount": 100000,
                "confirmations": 2,
            }
        }
    },
    ...
}
* // 
	
```

Transfer ltc
```
txId,err := w.Transfer("tb1q0r23g66m9rhhak8aahsg53wfp5egt2huuc4tnu",10000)
txId // string
```

### LTC Faucet
check this website https://testnet-faucet.com/ltc-testnet/

### Provider 
sign up and create free account at https://blockdaemon.com and use your own node  
```
config = ConfigBlockDaemon{
		Protocol: "litecoin",
		Network:  "mainnet",
		Token:    "your token here",
}
node := Node{
		Config: config,
		Test:   false, // or true if you want to use testnet network
	}
```
### Important
I simplified this repository github.com/ltcsuite/ltcd repository to create this package You can check go it for better examples and functionalities and do not use this package in production, I created this package for education purposes.


### Donation
Address `ltc1q5m8nzek2459kvqvnmshhluc56lgl0mjyc8p3dc`
