package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

type Instance struct {
	host string
	port int
}

func NewInstance(host string, port int) Instance {
	return Instance{
		host: host,
		port: port,
	}
}

type Balance struct {
	curIndex int
}

func (p *Balance) RoundRobinBalance(insts []Instance) (inst Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("no instance")
		return
	}
	lens := len(insts)
	if p.curIndex >= lens {
		p.curIndex = 0
	}
	inst = insts[p.curIndex]
	p.curIndex++
	return
}

func (p *Balance) RandomBalance(insts []Instance) (inst Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("no instance")
		return
	}
	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]
	return
}

func (i *Instance) readAndWritePort(c chan string) {
	select {
	case ports := <-c:
		newPorts := ports + strconv.Itoa(i.port)
		fmt.Println(newPorts)
		c <- newPorts
	default:
		//the first one
		fst := strconv.Itoa(i.port)
		fmt.Println(fst)
		c <- fst
	}
}
