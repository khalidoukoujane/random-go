package random

import "fmt"

func Comb()  {
	first := '0'
	second := '1'
	third := '2'
	for first <= '7' {
		second = first + 1
		for second <= '8' {
			third = second + 1
			for third <= '9' {

				fmt.Printf("%c",first)
				fmt.Printf("%c",second)
				fmt.Printf("%c",third)
				if first != '7' {
					fmt.Printf(", ")
				}
				third++
			}
			second++
		}
		first++
	}
}

//output

/*
$ go run main.go 
012, 013, 014, 015, 016, 017, 018, 019, 023, ..., 789
*/