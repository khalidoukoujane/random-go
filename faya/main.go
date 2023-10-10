package main

import (
	"fmt"
	"os"
)


func main() {
	FirstStr := os.Args[1:2]
	SecondStr := os.Args[2:3]
	s1 := []byte(FirstStr[0])
	s2 := []byte(SecondStr[0])
	Len := 0
	i := 0
	j := 0
	for i < len(s1){
		for j < len(s2){
			if s1[i] == s2[j]{
				Len++
				break
			}
			j++
		}
		i++
	}
	if Len == len(FirstStr[0]){
		fmt.Println(FirstStr[0])
	}
}