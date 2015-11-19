package main

import (
	"fmt"
	"github.com/zankich/concourse-scratch/client1"
	"github.com/zankich/concourse-scratch/client2"
)

func main() {
	fmt.Println(ReturnThree())
}

func ReturnThree() int {
	return client1.ReturnOne() + client2.ReturnTwo()
}
