package core

import "io"

type Transaction struct {
}

func (tx *Transaction) DecodedBinary(r io.Reader) error {
	return nil
}

func (tx *Transaction) EncodeBinary(w io.Writer) error {
	return nil
}
