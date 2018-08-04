package main

import (
	"errors"
	"fmt"
	"net"
	"runtime"

	"github.com/totoleo/yar-go/yar"
)

type Args struct {
	A, B int
	C    string
}

type Arith int

func init() {
}

func (t *Arith) Multiply(args *Args, reply *Args) error {
	reply.A = args.A * args.B
	reply.C = args.C + "_hello"
	return errors.New("hello")
}

func main() {
	var worker = runtime.NumCPU()
	runtime.GOMAXPROCS(worker)

	var server = yar.NewServer()
	server.Register(new(Arith))

	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println(err)
		return
	}
	server.Accept(listener)

}
