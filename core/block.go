package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/nerdishhomosapein/go-block-x/crypto"
	"github.com/nerdishhomosapein/go-block-x/types"
)

type Header struct {
	Version       uint32
	DataHash      types.Hash
	PrevBlockHash types.Hash
	TimeStamp     uint64
	Height        uint32
	Nonce         uint64
}

type Block struct {
	*Header
	Transactions []Transaction
	Validator    crypto.PublicKey
	Signature    *crypto.Signature

	// cached version of the header hash
	hash types.Hash
}

func NewBlock(h *Header, txx []Transaction) *Block {
	return &Block{
		Header:       h,
		Transactions: txx,
	}
}

func (b *Block) Sign(privKey crypto.PrivateKey) error {

	sig, err := privKey.Sign(b.HeaderByByte())
	if err != nil {
		return err
	}

	b.Validator = privKey.PublicKey()
	b.Signature = sig
	return nil

}

func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("Block has no signature")
	}

	if b.Signature.Verify(b.Validator, b.HeaderByByte()) {
		return fmt.Errorf("Block has invalid signature")
	}
	return nil
}

func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}
	return b.hash
}

func (b *Block) HeaderByByte() []byte {

	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	err := enc.Encode(b.Header)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}
