package random

import "fmt"

func PrintAlpha()  {
	c := 'a'
	for c < 'z'{
		fmt.Printf("%c",c)
		c++
	}
}