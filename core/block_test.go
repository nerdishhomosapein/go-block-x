package core

import (
	"bytes"
	"testing"
	"time"

	"github.com/nerdishhomosapein/go-block-x/types"
	"github.com/stretchr/testify/assert"
)

func TestHeader_Encode_Decode_Binary(t *testing.T) {

	h := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		TimeStamp: uint64(time.Now().UnixNano()),
		Height:    10,
		Nonce:     989394,
	}

	buf := &bytes.Buffer{}
	h.EncodeBinary(buf)
	assert.Nil(t, h.EncodeBinary(buf))

	hDecodec := &Header{}
	assert.Nil(t, hDecodec.DecodeBinary(buf))
	assert.Equal(t, hDecodec, h)
}

func TestBlock_Encode_Decode(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			TimeStamp: uint64(time.Now().UnixNano()),
			Height:    10,
			Nonce:     989394,
		},
		Transactions: nil,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buf))
	assert.Equal(t, b, bDecode)
}

func TestBlockHash(t *testing.T) {
	b := &Block{
		Header: Header{
			Version:   1,
			PrevBlock: types.RandomHash(),
			TimeStamp: uint64(time.Now().UnixNano()),
			Height:    10,
			Nonce:     989394,
		},
		Transactions: nil,
	}
	h := b.Hash()
	assert.False(t, h.IsZero())
}
