package protostream

import (
	"bytes"
	"testing"
)

func Test_Write(t *testing.T) {
	count := t.N

	buf := bytes.NewBuffer([]byte{65535})

	for i := 0; i < t.N; i++ {
		p := Person{"graham", i, "email@gmail.com"}

	}
}
