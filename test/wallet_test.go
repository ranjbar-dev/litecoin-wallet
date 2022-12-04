package test

import (
	litecoinWallet "github.com/ranjbar-dev/litecoin-wallet"
	"testing"
)

// GeneratelitecoinWallet test
func TestGenerateWallet(t *testing.T) {
	w := litecoinWallet.GenerateLitecoinWallet(node)
	if w == nil {
		t.Errorf("GeneratelitecoinWallet res was incorect, got: %q, want: %q.", "wallet", "*litecoinWallet")
	}
	if len(w.PrivateKey) == 0 {
		t.Errorf("GeneratelitecoinWallet PrivateKey was incorect, got: %q, want: %q.", w.PrivateKey, "valid PrivateKey")
	}
	if len(w.PublicKey) == 0 {
		t.Errorf("GeneratelitecoinWallet PublicKey was incorect, got: %q, want: %q.", w.PublicKey, "valid PublicKey")
	}
	if len(w.Address) == 0 {
		t.Errorf("GeneratelitecoinWallet Address was incorect, got: %q, want: %q.", w.Address, "valid Address")
	}
}

// CreatelitecoinWallet test
func TestCreateWallet(t *testing.T) {
	_, err := litecoinWallet.CreateLitecoinWallet(node, invalidPrivateKey)
	if err == nil {
		t.Errorf("CreatelitecoinWallet error was incorect, got: %q, want: %q.", err, "not nil")
	}

	w, err := litecoinWallet.CreateLitecoinWallet(node, validPrivateKey)
	if err != nil {
		t.Errorf("CreatelitecoinWallet error was incorect, got: %q, want: %q.", err, "nil")
	}
	if len(w.PrivateKey) == 0 {
		t.Errorf("CreatelitecoinWallet PrivateKey was incorect, got: %q, want: %q.", w.PrivateKey, "valid PrivateKey")
	}
	if len(w.PublicKey) == 0 {
		t.Errorf("CreatelitecoinWallet PublicKey was incorect, got: %q, want: %q.", w.PublicKey, "valid PublicKey")
	}
	if len(w.Address) == 0 {
		t.Errorf("CreatelitecoinWallet Address was incorect, got: %q, want: %q.", w.Address, "valid Address")
	}
	if len(w.Address) == 0 {
		t.Errorf("CreatelitecoinWallet AddressBase58 was incorect, got: %q, want: %q.", w.Address, "valid Address")
	}
}

// PrivateKeyRCDSA test
func TestPrivateKeyRCDSA(t *testing.T) {
	w := wallet()

	_, err := w.PrivateKeyRCDSA()
	if err != nil {
		t.Errorf("PrivateKeyRCDSA error was incorect, got: %q, want: %q.", err, "nil")
	}
}

// PrivateKeyBTCE test
func TestPrivateKeyBTCE(t *testing.T) {
	w := wallet()

	_, err := w.PrivateKeyBTCE()
	if err != nil {
		t.Errorf("PrivateKeyRCDSA error was incorect, got: %q, want: %q.", err, "nil")
	}
}

// PrivateKeyBytes test
func TestPrivateKeyBytes(t *testing.T) {
	w := wallet()

	bytes, err := w.PrivateKeyBytes()
	if err != nil {
		t.Errorf("PrivateKeyBytes error was incorect, got: %q, want: %q.", err, "nil")
	}
	if len(bytes) == 0 {
		t.Errorf("PrivateKeyBytes bytes len was incorect, got: %q, want: %q.", len(bytes), "more than 0")
	}
}

// Balance test
func TestBalance(t *testing.T) {
	w := wallet()

	_, err := w.Balance()
	if err != nil {
		t.Errorf("Balance error was incorect, got: %q, want: %q.", err, "nil")
	}
}

// Transfer test
func TestTransfer(t *testing.T) {
	w := wallet()

	txId, err := w.Transfer(validToAddress, ltcAmount)
	if err != nil {
		t.Errorf("Transfer error was incorect, got: %q, want: %q.", err, "nil")
	}
	if len(txId) == 0 {
		t.Errorf("Transfer txId was incorect, got: %q, want: %q.", txId, "not nil")
	}
}

// TestEstimateTransfer test
func TestEstimateTransfer(t *testing.T) {
	w := wallet()

	amount, err := w.EstimateTransferFee(validToAddress, ltcAmount)

	if err != nil {
		t.Errorf("EstimateTransferFee error was incorect, got: %q, want: %q.", err, "nil")
	}
	if amount == 0 {
		t.Errorf("EstimateTransferFee amount was incorect, got: %q, want: %q.", amount, "not 0")
	}
}
