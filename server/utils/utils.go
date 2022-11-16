package utils

func MapSpace(str []rune) []rune {
	ret := []rune{}
	for i, l := range str {
		if i < len(str)-1 && (str[i+1] == '_' || l == '_') {
			ret = append(ret, l, ' ')
		} else {
			ret = append(ret, l)
		}
	}

	return ret
}
