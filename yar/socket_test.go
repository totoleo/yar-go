package yar

import (
	"fmt"
	"testing"
)

func TestDialSocket(t *testing.T) {
	var cli, err = Dial("tcp", ":12345", "msgpack") //msgpack json
	if err != nil {
		panic(err)
	}

	var args = map[string]interface{}{"A": 4, "B": 5}
	var reply = make(map[string]interface{})
	fmt.Println("goroutine start...")
	if err := cli.Call("Arith.Multiply", args, &reply); err != nil {
		t.Error(err)
		return
	}
	fmt.Println(reply)
	t.Log(reply)
}
