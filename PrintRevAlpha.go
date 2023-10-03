package random

import "fmt"

func PrintRevAlpha()  {
	c := 'z'
	for c > 'a'{
		fmt.Printf("%c",c)
		c--
	}
}