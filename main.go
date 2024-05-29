package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// read text function
	rt := func(f *os.File) {
		s, er := ioutil.ReadAll(f)
		if er != nil {
			panic(er)
		}
		fmt.Println(string(s))
	}

	fn := "data.txt"

	f, er := os.OpenFile(fn, os.O_RDONLY, os.ModePerm)
	if er != nil {
		panic(er)
	}

	fmt.Println("<<< start >>>")
	rt(f)
	fmt.Println("<<< end >>>")

	defer f.Close()
}
