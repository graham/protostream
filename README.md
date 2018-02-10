# ProtoStream
A simple way to serialize your protobufs to and from io.Reader/io.Writer. Some benchmarks to show speed differences.

## Usage
The package provides:

    - Write(io.Writer, proto.Message)
      Write protobuf to writer.

    - Read(io.Reader, proto.Message)
      Read message into protobuf.

    - WriteFromChan(io.Writer, chan proto.Message)
      Continously write protobufs to writer from a channel.

    - ReadToChan(io.Reader, proto.Message, chan proto.Message)
      Read a reader and unmarshal protobufs onto a channel.

Channel methods are slower, but make life easier. The performance is very nice compared to something like JSON (which is actually quite fast in Go).

## Benchmarks
You'll need to compile the included protobuf in order to run benchmarks:

```protoc --gofast_out=import_path=protostream:. tutorial.proto```

An example run:

```
goos: darwin
goarch: amd64
pkg: github.com/graham/protostream
Benchmark_PureWrite-8        	10000000	       206 ns/op
Benchmark_PureRead-8         	10000000	       200 ns/op
Benchmark_ChanRead-8         	 1000000	      1569 ns/op
Benchmark_ChanWrite-8        	 3000000	       417 ns/op
Benchmark_PureWrite_json-8   	 3000000	       579 ns/op
Benchmark_PureRead_json-8    	 1000000	      1626 ns/op
PASS
ok  	github.com/graham/protostream	14.401s
```

Author: Graham Abbott