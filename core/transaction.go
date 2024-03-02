package core

import (
	"fmt"

	"github.com/nerdishhomosapein/go-block-x/crypto"
)

type Transaction struct {
	Data []byte

	PublicKey crypto.PublicKey
	signature *crypto.Signature
}

func (tx *Transaction) Sign(privKey crypto.PrivateKey) error {

	sig, err := privKey.Sign(tx.Data)
	if err != nil {
		return err
	}

	tx.PublicKey = privKey.PublicKey()
	tx.signature = sig
	return nil
}

func (tx *Transaction) Verify() error {

	if tx.signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	if !tx.signature.Verify(tx.PublicKey, tx.Data) {
		return fmt.Errorf("Invalid Transaction signature")
	}

	return nil
}
