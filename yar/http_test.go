package yar

import (
	"fmt"
	"testing"
)

func TestDialHTTP(t *testing.T) {
	var cli, err = DialHTTP("http://zee.test:80", "msgpack") //msgpack json
	if err != nil {
		panic(err)
	}

	var args = map[string]interface{}{"a": 4}
	var reply = new(int64)
	fmt.Println("goroutine start...")
	if err := cli.Call("add", args, reply); err != nil {
		t.Error(err)
		return
	}
	fmt.Println(*reply)
}

/*
func testClient(wg *sync.WaitGroup) {
	t1 := time.Now()
	var client, err = yar.Dial("tcp", "192.168.125.185:12345", "json") //msgpack json
	if err != nil {
		fmt.Println("Dialing:", err)
	}
	for i := 0; i < 1000; i++ {
		var args = &Args{4, 5, "GO"}
		var reply = &Args{}

		err := client.Call("Arith.Multiply", args, &reply)
		if err != nil {
			fmt.Println("Call:", err)
		}

		//fmt.Printf("Result: %d * %d = %d", args.A, args.B, reply)
	}
	t2 := time.Now()
	fmt.Println(" lasted ", t2.Sub(t1))
	wg.Done()
} */
