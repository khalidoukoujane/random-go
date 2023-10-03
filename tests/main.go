package main

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

func IsNegative(num int)  {
	if num > 0 {
		fmt.Println("P")
	} else {
		fmt.Println("N")
	}
}

func PrintAlpha()  {
	c := 'a'
	for c < 'z'{
		fmt.Printf("%c",c)
		c++
	}
}

func PrintNums()  {
	c := '0'
	for c < '9'{
		fmt.Printf("%c",c)
		c++
	}
}

func PrintRevAlpha()  {
	c := 'z'
	for c > 'a'{
		fmt.Printf("%c",c)
		c--
	}
}

func main()  {
	Comb()
	IsNegative(2)
	PrintAlpha()
	PrintNums()
	PrintRevAlpha()
	fmt.Printf("\n")
}