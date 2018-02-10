package protostream

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"testing"

	"github.com/gogo/protobuf/proto"
)

func GenReader(b *testing.B) *bytes.Reader {
	buf := bytes.NewBuffer([]byte{})
	t := &Person{
		Name:  "Graham",
		Id:    123,
		Email: "graham.abbott@gmail.com",
	}

	for i := 0; i < b.N; i++ {
		Write(buf, t)
	}

	return bytes.NewReader(buf.Bytes())
}

func Benchmark_PureWrite(b *testing.B) {
	GenReader(b)
}

func Benchmark_PureRead(b *testing.B) {
	buf := GenReader(b)
	pb := Person{}

	b.ResetTimer()

	for {
		_, err := Read(buf, &pb)

		if pb.Id != 123 {
			panic("Not equal")
		}

		if err != nil {
			return
		}
	}
}

func Benchmark_ChanRead(b *testing.B) {
	buf := GenReader(b)
	pb := Person{}
	c := make(chan proto.Message, 1)
	go ReadToChan(buf, &pb, c)

	b.ResetTimer()
	for p := range c {
		if p.(*Person).Id != 123 {
			panic("not equal")
		}
	}
}

func Benchmark_ChanWrite(b *testing.B) {
	buf := bytes.NewBuffer([]byte{})
	c := make(chan proto.Message, 1)

	t := &Person{
		Name:  "Graham",
		Id:    123,
		Email: "graham.abbott@gmail.com",
	}
	go WriteFromChan(buf, c)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c <- t
	}

}

type PersonJson struct {
	Name  string
	Id    int
	Email string
}

func GenReader_json(b *testing.B) *bytes.Reader {
	buf := bytes.NewBuffer([]byte{})
	t := PersonJson{"Graham", 123, "graham.abbott@gmail.com"}
	enc := json.NewEncoder(buf)

	for i := 0; i < b.N; i++ {
		enc.Encode(t)
	}

	return bytes.NewReader(buf.Bytes())
}

func Benchmark_PureWrite_json(b *testing.B) {
	GenReader_json(b)
}

func Benchmark_PureRead_json(b *testing.B) {
	buf := GenReader_json(b)
	dec := json.NewDecoder(buf)

	var p PersonJson

	b.ResetTimer()

	for {
		if err := dec.Decode(&p); err == io.EOF {
			return
		} else if err != nil {
			log.Fatal(err)
		}
		if p.Id != 123 {
			panic("not equal")
		}
	}
}
