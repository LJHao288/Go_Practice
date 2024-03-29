package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//Generate Instance Slice
	var insts []Instance
	for i := 0; i < 10; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		port, _ := strconv.Atoi(fmt.Sprintf("880%d", i))
		one := NewInstance(host, port)
		insts = append(insts, one)
	}

	var portsChannel = make(chan string)

	//load balance
	lb := Balance{curIndex: 0}
	for {
		inst, err := lb.RoundRobinBalance(insts)
		if err != nil {
			fmt.Println("do balance err")
			time.Sleep(time.Second)
			continue
		}
		fmt.Println(inst)
		go inst.readAndWritePort(portsChannel)
		time.Sleep(time.Second)
	}

} //main func
