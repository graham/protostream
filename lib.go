package protostream

import (
	"encoding/binary"
	"io"

	"github.com/gogo/protobuf/proto"
)

type ProtoWriter struct {
	io.Writer
}

const BYTES_FOR_SIZE = 8

func Write(w io.Writer, mesg proto.Message) (n int, err error) {
	b, err := proto.Marshal(mesg)

	if err != nil {
		return n, err
	}

	messageSize := len(b)
	sizeBuf := make([]byte, BYTES_FOR_SIZE)
	binary.BigEndian.PutUint64(sizeBuf, uint64(messageSize))

	var total int = BYTES_FOR_SIZE

	n, err = w.Write(sizeBuf)
	total += n
	if err != nil {
		return n, err
	}

	n, err = w.Write(b)
	total += n

	return total, err
}

func Read(r io.Reader, pb proto.Message) (n int, err error) {
	sizeBuf := make([]byte, BYTES_FOR_SIZE)

	_, err = io.ReadFull(r, sizeBuf)
	if err != nil {
		return n, err
	}

	var total int = BYTES_FOR_SIZE

	messageSize := binary.BigEndian.Uint64(sizeBuf)
	message := make([]byte, messageSize)

	n, err = io.ReadFull(r, message)
	total += n

	if err != nil || n == 0 {
		return n, err
	}

	return total, proto.Unmarshal(message, pb)
}

func WriteFromChan(w io.Writer, ch chan proto.Message) {
	for msg := range ch {
		Write(w, msg)
	}
}

func ReadToChan(r io.Reader, pb proto.Message, ch chan proto.Message) {
	var fileOffset int64 = 0

	for {
		item := proto.Clone(pb)
		n, err := Read(r, item)

		fileOffset += int64(n)

		if err == io.EOF {
			if n > 0 {
				ch <- item
			}
			close(ch)
			return
		}

		if err != nil {
			panic(err)
		}

		ch <- item
	}
}
