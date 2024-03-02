package core

import (
	"crypto/sha256"

	"github.com/nerdishhomosapein/go-block-x/types"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlockHasher struct {
}

func (BlockHasher) Hash(b *Block) types.Hash {

	h := sha256.Sum256(b.HeaderByByte())
	return types.Hash(h)
}
