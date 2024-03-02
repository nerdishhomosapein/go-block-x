package core

import (
	"testing"

	"github.com/nerdishhomosapein/go-block-x/crypto"
	"github.com/stretchr/testify/assert"
)

func TestSignTransaction(t *testing.T) {
	data := []byte("foo")
	privKey := crypto.GeneratePrivatekey()
	tx := &Transaction{
		Data: data,
	}
	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.signature)
}

func TestVerfifyTransaction(t *testing.T) {

	privKey := crypto.GeneratePrivatekey()
	tx := &Transaction{
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.Nil(t, tx.Verify())

	otherPrivKey := crypto.GeneratePrivatekey()
	tx.PublicKey = otherPrivKey.PublicKey()

	assert.NotNil(t, tx.Verify())
}
