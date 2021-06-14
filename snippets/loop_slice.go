package main

import (
	"fmt"
	"strings"
)

func main(){
	s := "The quick brown fox jumps over the lazy dog"
	sl := strings.Split(s," ")
	for _, i := range(sl) {
		fmt.Println(i)
	}
}
