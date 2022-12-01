package litecoinWallet

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ltcsuite/ltcd/chaincfg"
	"github.com/ltcsuite/ltcd/ltcutil"
	"github.com/ranjbar-dev/litecoin-wallet/blockDaemon"
	"github.com/ranjbar-dev/litecoin-wallet/blockDaemon/response"
	"github.com/ranjbar-dev/litecoin-wallet/enums"
	"strconv"
)

type LitecoinWallet struct {
	Node       enums.Node
	Address    string
	PrivateKey string
	PublicKey  string
	bd         blockDaemon.BlockDaemon
}

// generating

func GenerateLitecoinWallet(node enums.Node) *LitecoinWallet {

	privateKey, _ := generatePrivateKey()
	privateKeyHex := convertPrivateKeyToHex(privateKey)

	publicKey, _ := getPublicKeyFromPrivateKey(privateKey)
	publicKeyHex := convertPublicKeyToHex(publicKey)

	address, _ := getAddressFromPrivateKey(node, privateKey)

	return &LitecoinWallet{
		Node:       node,
		Address:    address,
		PrivateKey: privateKeyHex,
		PublicKey:  publicKeyHex,
		bd:         blockDaemon.NewBlockDaemonService(node.Config),
	}
}

func CreateLitecoinWallet(node enums.Node, privateKeyHex string) (*LitecoinWallet, error) {

	privateKey, err := privateKeyFromHex(privateKeyHex)
	if err != nil {
		return nil, err
	}

	publicKey, err := getPublicKeyFromPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	publicKeyHex := convertPublicKeyToHex(publicKey)

	address, err := getAddressFromPrivateKey(node, privateKey)
	if err != nil {
		return nil, err
	}

	return &LitecoinWallet{
		Node:       node,
		Address:    address,
		PrivateKey: privateKeyHex,
		PublicKey:  publicKeyHex,
		bd:         blockDaemon.NewBlockDaemonService(node.Config),
	}, nil
}

// struct functions

func (bw *LitecoinWallet) Chain() *chaincfg.Params {
	chainConfig := &chaincfg.MainNetParams
	if bw.Node.Test {
		chainConfig = &chaincfg.TestNet4Params
	}
	return chainConfig
}

func (bw *LitecoinWallet) PrivateKeyRCDSA() (*ecdsa.PrivateKey, error) {
	return privateKeyFromHex(bw.PrivateKey)
}

func (bw *LitecoinWallet) PrivateKeyBTCE() (*btcec.PrivateKey, error) {

	temp, err := bw.PrivateKeyBytes()
	if err != nil {
		return nil, err
	}

	priv, _ := btcec.PrivKeyFromBytes(temp)

	return priv, nil
}

func (bw *LitecoinWallet) PrivateKeyBytes() ([]byte, error) {

	priv, err := bw.PrivateKeyRCDSA()
	if err != nil {
		return []byte{}, err
	}

	return crypto.FromECDSA(priv), nil
}

func (bw *LitecoinWallet) WIF() (*ltcutil.WIF, error) {

	priv, err := bw.PrivateKeyBTCE()
	if err != nil {
		return nil, err
	}

	return ltcutil.NewWIF(priv, bw.Chain(), true)
}

// private key

func generatePrivateKey() (*ecdsa.PrivateKey, error) {

	return crypto.GenerateKey()
}

func convertPrivateKeyToHex(privateKey *ecdsa.PrivateKey) string {

	privateKeyBytes := crypto.FromECDSA(privateKey)

	return hexutil.Encode(privateKeyBytes)[2:]
}

func privateKeyFromHex(hex string) (*ecdsa.PrivateKey, error) {

	return crypto.HexToECDSA(hex)
}

// public key

func getPublicKeyFromPrivateKey(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, error) {

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error in getting public key")
	}

	return publicKeyECDSA, nil
}

func convertPublicKeyToHex(publicKey *ecdsa.PublicKey) string {

	privateKeyBytes := crypto.FromECDSAPub(publicKey)

	return hexutil.Encode(privateKeyBytes)[2:]
}

// address

func getAddressFromPrivateKey(node enums.Node, privateKey *ecdsa.PrivateKey) (string, error) {

	chainConfig := &chaincfg.MainNetParams
	if node.Test {
		chainConfig = &chaincfg.TestNet4Params
	}

	_, pub := btcec.PrivKeyFromBytes(crypto.FromECDSA(privateKey))

	addr, err := ltcutil.NewAddressWitnessPubKeyHash(ltcutil.Hash160(pub.SerializeCompressed()), chainConfig)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return addr.EncodeAddress(), nil
}

// balance

func (bw *LitecoinWallet) Balance() (int64, error) {

	res, err := bw.bd.AddressBalance(bw.Address)
	if err != nil {
		return 0, err
	}

	balance, err := strconv.Atoi(res[0].ConfirmedBalance)
	if err != nil {
		return 0, err
	}

	return int64(balance), nil
}

// transactions

func (bw *LitecoinWallet) UTXOs() ([]response.UTXO, error) {

	var res []response.UTXO

	utxos, err := bw.bd.AddressUTXO(bw.Address)
	if err != nil {
		return nil, err
	}

	for _, utxo := range utxos.Data {
		if utxo.Mined.Confirmations > 2 {
			res = append(res, utxo)
		}
	}

	return res, nil
}

func (bw *LitecoinWallet) Txs() ([]response.Transaction, error) {

	res, err := bw.bd.AddressTxs(bw.Address)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func (bw *LitecoinWallet) Transfer(toAddress string, amountInLitoshi int64) (string, error) {

	privateKey, err := bw.PrivateKeyBTCE()
	if err != nil {
		return "", err
	}

	return createSignAndBroadcastTransaction(bw.Chain(), privateKey, bw.Address, toAddress, amountInLitoshi)
}
