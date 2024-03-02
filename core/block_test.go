package core

import (
	"testing"
	"time"

	"github.com/nerdishhomosapein/go-block-x/crypto"
	"github.com/nerdishhomosapein/go-block-x/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		TimeStamp:     uint64(time.Now().UnixNano()),
	}

	tx := Transaction{
		Data: []byte("foo"),
	}
	return NewBlock(header, []Transaction{tx})
}

func TestSignBlock(t *testing.T) {

	privKey := crypto.GeneratePrivatekey()
	b := randomBlock(0)
	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
}

// func TestVerifyBlock(t *testing.T) {
//
// 	privKey := crypto.GeneratePrivatekey()
// 	b := randomBlock(0)
//
// 	assert.Nil(t, b.Sign(privKey))
// 	// assert.Nil(t, b.Verify())
//
// 	otherPrivKey := crypto.GeneratePrivatekey()
// 	b.Validator = otherPrivKey.PublicKey()
// 	assert.NotNil(t, b.Verify())
//
// 	b.Height = 100
// 	assert.NotNil(t, b.Verify())
// }
