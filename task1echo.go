package main

import (
	"fmt"
	"os"
	"strings"
)

func echo() {
	echo1()
	echo2()
	echo3()
}

func echo1() {
	var s, sep string

	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(os.Args[0])
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "\n"
	}
	fmt.Println(os.Args[0])
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[0:], "\n"))
}