package fuzzPgxmock

import "testing"

func FuzzPgxmock(f *testing.F) {
	f.Add([]byte("hello"))
	f.Fuzz(func(t *testing.T, data []byte) {
		mayhemit(data)
	})
}
