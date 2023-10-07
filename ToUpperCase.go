package random

func ToUpperCase(s string) string {
	str := []rune(s)
	l := len(str)
	i := 0
	var res []rune
	for i < l{
		
		res = append(res,rune(str[i] - ' '))
		i++
	}
	return(string(res))
}
