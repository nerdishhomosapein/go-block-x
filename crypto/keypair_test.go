package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {

	privKey := GeneratePrivatekey()
	pubKey := privKey.PublicKey()
	// address := pubKey.Address()

	msg := []byte("hello, world!!!")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)
	assert.True(t, sig.Verify(pubKey, msg))

}
