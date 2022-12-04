package litecoinWallet

import (
	"errors"
	"github.com/ltcsuite/ltcd/chaincfg"
	"github.com/ltcsuite/ltcd/chaincfg/chainhash"
	"github.com/ltcsuite/ltcd/ltcutil"
	"github.com/ltcsuite/ltcd/txscript"
	"github.com/ltcsuite/ltcd/wire"
	"github.com/ranjbar-dev/litecoin-wallet/blockDaemon"
	"github.com/ranjbar-dev/litecoin-wallet/enums"
)

func estimateTransactionFee(chain *chaincfg.Params, fromAddress string, toAddress string, amount int64) (int64, error) {

	node := enums.MAIN_NODE
	if &chaincfg.TestNet4Params == chain {
		node = enums.TEST_NODE
	}

	toAddr, err := ltcutil.DecodeAddress(toAddress, chain)
	if err != nil {
		return 0, errors.New("DecodeAddress destAddrStr err " + err.Error())
	}

	toAddressByte, err := txscript.PayToAddrScript(toAddr)
	if err != nil {
		return 0, errors.New("toAddr PayToAddrScript err " + err.Error())
	}

	bd := blockDaemon.NewBlockDaemonService(node.Config)

	res, err := bd.EstimateFee()
	if err != nil {
		return 0, err
	}

	utxoList, _, err := prepareUTXOForTransaction(chain, fromAddress, amount)
	if err != nil {
		return 0, errors.New("vin err " + err.Error())
	}
	if len(utxoList) == 0 {
		return 0, errors.New("insufficient balance")
	}

	transaction := wire.NewMsgTx(2)

	for _, utxo := range utxoList {

		hash, err := chainhash.NewHashFromStr(utxo.Mined.TxId)
		if err != nil {
			return 0, err
		}

		txIn := wire.NewTxIn(wire.NewOutPoint(hash, uint32(utxo.Mined.Index)), nil, [][]byte{})
		txIn.Sequence = txIn.Sequence - 2
		transaction.AddTxIn(txIn)
	}

	transaction.AddTxOut(wire.NewTxOut(amount, toAddressByte))

	return int64(transaction.SerializeSize() * res.EstimatedFees.Slow), nil
}
