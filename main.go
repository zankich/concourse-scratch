package main

import (
	"fmt"
	"github.com/zankich/concourse-scratch/client1"
	"github.com/zankich/concourse-scratch/client2"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 2 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
		if buf, err := ioutil.ReadFile("version"); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(buf))
		}
		os.Exit(0)
	}
	fmt.Printf("The answer is: %d!\n", ReturnThree())
}

func ReturnThree() int {
	return client1.ReturnOne() + client2.ReturnTwo()
}
