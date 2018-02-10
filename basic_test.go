package protostream

import (
	"bytes"
	"testing"
)

func Test_Write(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	var p Person

	var count int = 0
	var expectedSize = 0
	var iterSize int = 10

	for i := 0; i < iterSize; i++ {
		p = Person{"graham", int32(i + iterSize), "email@gmail.com"}
		n, _ := Write(buf, &p)
		count += n
		expectedSize += p.Size()
	}

	if count != expectedSize+(BYTES_FOR_SIZE*iterSize) {
		t.Fail()
	}
}
